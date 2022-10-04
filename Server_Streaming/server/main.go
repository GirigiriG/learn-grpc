package main

import (
	"book/handler"
	"book/proto/book"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	tcpServer, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Printf("failed to connect to port 9090 %v", err)
	}

	// Create new GRPC server
	server := grpc.NewServer()
	
	book.RegisterBookServiceServer(server, &handler.Server{})

	if err = server.Serve(tcpServer); err != nil {
		log.Fatalf("Failed to start GRCP server %v", err)
	}

}
