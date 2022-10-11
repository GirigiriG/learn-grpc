package handler

import (
	"event/proto/event"
	"io"
	"log"
)

type Server struct {
	event.UnimplementedEventServiceServer
}

func (s *Server) FullDuplexStream(stream event.EventService_FullDuplexStreamServer) error {
	log.Println("Proccessing client stream from the server")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("An occured while reading client stream %v", err)
		}

		res := "Hello " + req.Title

		err = stream.Send(&event.Response{Result: res})
		if err != nil {
			log.Fatalf("Error occurred while sending message to the client %v", err)
		}

	}
}
