package main

import (
	"context"
	"log"
	"net"

	pb "github.com/GirigiriG/learn-grpc/greet/proto/greet/proto"
	"google.golang.org/grpc"
)

const (
	GRPCServerAddress = "0.0.0.0:5051"
)

type GreetServer struct {
	pb.UnimplementedGreetServiceServer
}

func (s *GreetServer) sayHello(ctx context.Context, in *pb.GreetRequest) *pb.GreetResponse {
	log.Printf("recieved %v", in.GetFirstName())
	var name string = "Gideon from grpc"
	return &pb.GreetResponse{Result: name}
}

func main() {
	grpcServerListner, err := net.Listen("tcp", GRPCServerAddress)
	if err != nil {
		log.Printf("Failed to connect to grpc server %v \n", GRPCServerAddress)
		return
	}

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &GreetServer{})
	

	//start the server

	s.Serve(grpcServerListner)
}
