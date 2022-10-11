// Events client
package main

import (
	"context"
	"event/proto/event"
	"io"
	"log"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error occured while trying to connect to event service %v", err)
	}

	defer conn.Close()

	c := event.NewEventServiceClient(conn)
	callDuplexStream(c)
}

func callDuplexStream(c event.EventServiceClient) {
	stream, err := c.FullDuplexStream(context.Background())
	if err != nil {
		log.Fatalf("Error occured calling the FullDuplexStream RPC")
	}

	reqs := []*event.Request{
		{Title: "Saving DevOps one Config At A Time "},
		{Title: ""},
		{Title: "Test one from client"},
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for _, req := range reqs {
			stream.Send(req)
		}

		stream.CloseSend()
	}()

	go func() {
		defer wg.Done()
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error reading from the server's stream %v", err)
				break
			}

			log.Printf("Message from server: %s\n", res.Result)
		}

	}()
	wg.Wait()
}
