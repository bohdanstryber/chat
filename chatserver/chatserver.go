package chat

import (
	"log"
	"math/rand"
	"sync"
	"time"

	pb "github.com/bohdanstryber/chat"
)

type messageUnit struct {
	ClientName        string
	MessageBody       string
	MessageUniqueCode int
	ClientUniqueCode  int
}

type messageQue struct {
	MQue []messageUnit
	mu   sync.Mutex
}

var messageQueObject = messageQue{}

type ChatServerStruct struct {
	pb.UnimplementedChatServer
}

func (cs *ChatServerStruct) SendMessage(csi pb.Chat_SendMessageServer) error {
	clientUniqueCode := rand.Intn(1e3)

	go receiveFromStream(csi, clientUniqueCode)

	errCh := make(chan error)

	go sendToStream(csi, clientUniqueCode, errCh)

	return <-errCh
}

func receiveFromStream(cs pb.Chat_SendMessageServer, ClientUniqueCode int) {
	for {
		req, err := cs.Recv()

		if err != nil {
			log.Printf("Error reciving request from client :: %v", err)

			break
		} else {
			messageQueObject.mu.Lock()
			messageQueObject.MQue = append(messageQueObject.MQue, messageUnit{ClientName: req.Name, MessageBody: req.Body})
			messageQueObject.mu.Unlock()

			log.Printf("%v", messageQueObject.MQue[len(messageQueObject.MQue)-1])
		}
	}
}

func sendToStream(cs pb.Chat_SendMessageServer, clientUniqueCode int, errCh chan error) {
	for {
		for {
			time.Sleep(500 * time.Millisecond)
			messageQueObject.mu.Lock()

			if len(messageQueObject.MQue) == 0 {
				messageQueObject.mu.Unlock()

				break
			}

			senderUniqueCode := messageQueObject.MQue[0].ClientUniqueCode
			senderNameClient := messageQueObject.MQue[0].ClientName
			messageClient := messageQueObject.MQue[0].MessageBody

			messageQueObject.mu.Unlock()

			if senderUniqueCode != clientUniqueCode {
				err := cs.Send(&pb.FromServer{Name: senderNameClient, Body: messageClient})

				if err != nil {
					errCh <- err
				}

				messageQueObject.mu.Lock()

				if len(messageQueObject.MQue) >= 2 {
					messageQueObject.MQue = messageQueObject.MQue[1:]
				} else {
					messageQueObject.MQue = []messageUnit{}
				}

				messageQueObject.mu.Unlock()
			}
		}

		time.Sleep(1 * time.Second)
	}
}
