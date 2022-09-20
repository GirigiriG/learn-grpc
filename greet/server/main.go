package main

import (
	"log"
	"net"

	pb "github.com/GirigiriG/learn-grpc/greet/proto"
)

const (
	GRPCServerAddress = "0.0.0.0:5051"
)

type server struct {
	pb.GreetServiceServer
}

func main() {
	listen, err := net.Listen("tcp", GRPCServerAddress)
	if err != nil {
		log.Printf("Failed to connect to grpc server %v \n", GRPCServerAddress)
		return
	}

}
