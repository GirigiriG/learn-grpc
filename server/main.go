//server
package main

import (
	"log"
	"net"
	"profile/handler"
	"profile/proto/profile"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen to grpc %v", err)
	}

	server := grpc.NewServer()
	profile.RegisterProfileServiceServer(server, &handler.Profile{})

	if err = server.Serve(lis); err != nil {
		log.Fatalf("failed to start GRPC server %v", err)
	}

}
