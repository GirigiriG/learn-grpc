package main

import (
	"book/proto/book"
	"context"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	conn, err := grpc.Dial(":9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connect to Book service on port 9090 %v", err)
	}
	defer conn.Close()

	bookService := book.NewBookServiceClient(conn)

	readStreamOnBookCreation(bookService)
}

func readStreamOnBookCreation(bookService book.BookServiceClient) {
	stream, err := bookService.CreateBook(context.Background(), &book.BookRequest{
		Title:     "The art of writing good code",
		Author:    "Gideon Girigiri",
		CreatedAt: timestamppb.Now(),
	})
	if err != nil {
		log.Fatalf("Error calling book service %v", err)
	}

	for {
		// Read response from server streem
		serverStream, err := stream.Recv()
		// When the stream is closed or done sendig data it will be closed and returns an EOF
		if err == io.EOF {
			log.Println("Done reading....")
			break
		}

		if err != nil {
			log.Fatalf("Error read bookService create book stream %v", err)
		}

		log.Print(serverStream.Result.CreatedAt.AsTime())
	}
}
