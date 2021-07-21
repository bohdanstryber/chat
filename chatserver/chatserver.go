package main

import "sync"

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

type chatServer struct {
}

// func (cs *chatServer) chat(csi Chat_SendMessageServer) error {
//
// }

func receiveFromStream(csi grpcChatServer.chat_SendMessageServer, ClientUniqueCode int) {
	for {

	}
}
