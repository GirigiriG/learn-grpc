//Event handler
package handler

import (
	pb "event/proto/event"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
)

type Server struct {
	pb.UnimplementedEventServiceServer
}

func (s *Server) FullDuplexStream(stream pb.EventService_FullDuplexStreamServer) error {
	log.Println("Proccessing client stream from the server")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		} 

		if err != nil {
			log.Fatalf("An occured while reading client stream %v", err)
		}

		if req.Title == "" {
			return status.Error(codes.InvalidArgument, "Title can not be null")
		}
		res := "Hello " + req.Title

		err = stream.Send(&pb.Response{Result: res})
		if err != nil {
			log.Fatalf("Error occurred while sending message to the client %v", err)
		}
	}
}