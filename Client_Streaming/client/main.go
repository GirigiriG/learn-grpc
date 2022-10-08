package main

import (
	"context"
	"log"
	"logger/proto/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:9090"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connecting to logger service RPC server %v", err)
	}

	defer conn.Close()

	c := logger.NewLoggerServiceClient(conn)

	sendClientStream(c)
}

func sendClientStream(c logger.LoggerServiceClient) {

	reqs := []*logger.Request{
		{Message: "There was an successful update of record _id:19s43292", Level: "Info"},
		{Message: "Memory usage 70%", Level: "Warning"},
		{Message: "Error occurred unable to connnect to payment service", Level: "Fatal"},
	}

	stream, err := c.Streamer(context.Background())
	if err != nil {
		log.Fatalf("Error calling streaming RPC from logger service %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while recieving response from server %v", err)
	}

	log.Printf("%s\n", res.Result)
}
