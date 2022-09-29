// client
package main

import (
	"context"
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

}
