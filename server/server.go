package main

import (
	"log"
	"net"
	"os"

	pb "github.com/bohdanstryber/chat"
	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	listen, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Could not listen @ %v :: %v", port, err)
	}

	log.Panicln("Listening @ " + port)

	grpcServer := grpc.NewServer()

	err = grpcServer.Serve(listen)

	if err != nil {
		log.Fatalf("Failed to start gRPC server :: %v", err)
	}

	cs := pb.ChatServerStruct{}

	pb.RegisterChatServer(grpcServer, cs)
}
