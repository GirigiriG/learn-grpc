// Events server
package main

import (
	"event/handler"
	"event/proto/event"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Error occurred while listening on port 9090 %v", err)
	}

	server := grpc.NewServer()

	event.RegisterEventServiceServer(server, &handler.Server{})

	err = server.Serve(lis)

	if err != nil {
		log.Fatalf("Error occurred while start server %v", err)
	}
}
