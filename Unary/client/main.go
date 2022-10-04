// client
package main

import (
	"context"
	"io"
	"log"
	"profile/proto/profile"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	con, err := grpc.Dial(":9090", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to listen to grpc %v", err)
	}

	defer con.Close()

	c := profile.NewProfileServiceClient(con)

	req := profile.CreateRequest{Name: "Gideon Girigiri", Id: 123, IsValid: true}
	response, err := c.Create(context.Background(), &req)
	if err != nil {
		log.Fatalf("failed to create profile %v", err)
	}

	log.Println(response.Message)

	getProfileStream(c)
}

// Read stream from server
func getProfileStream(c profile.ProfileServiceClient) {
	stream, err := c.CreateProfileStream(context.Background(), &profile.CreateRequest{
		Name:    "Gideon",
		Id:      123,
		IsValid: false,
	})

	if err != nil {
		log.Fatalf("error while calling the streaming api %v\n", err)
	}

	// Loop and read server side stream until EOF
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			log.Println("Done reading")
			break
		}

		if err != nil {
			log.Fatalf("error occured while reading from stream %v", err)
		}

		log.Printf("message from the stream %v\n", msg.Message)
	}
}
