// server
package main

import (
	"log"
	"logger/handler"
	"logger/proto/logger"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen to grpc %v", err)
	}

	server := grpc.NewServer()
	logger.RegisterLoggerServiceServer(server, &handler.Server{})

	if err = server.Serve(lis); err != nil {
		log.Fatalf("failed to start GRPC server %v", err)
	}
}
