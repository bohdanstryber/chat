package main

import (
	"log"
	"net"
	"os"

	"grpcChatServer/chat"

	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	listen, err := net.Listen("tcp", ":"+port)

	if err != nil {
		log.Fatalf("Could not listen @ %v :: %v", port, err)
	}

	log.Println("Listening @ " + port)

	grpcServer := grpc.NewServer()

	err = grpcServer.Serve(listen)

	if err != nil {
		log.Fatalf("Failed to start gRPC server :: %v", err)
	}

	cs := chat.ChatServerStruct{}

	chat.RegisterChatServer(grpcServer, &cs)

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to start gRPC Server :: %v", err)
	}
}
