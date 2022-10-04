package handler

import (
	"book/proto/book"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	book.UnimplementedBookServiceServer
}

func (s *Server) CreateBook(req *book.BookRequest, stream book.BookService_CreateBookServer) error {
	for {
		time.Sleep(time.Millisecond * 200)
		stream.Send(&book.BookResponse{Result: &book.BookRequest{
			Title:     req.Title,
			Author:    req.Author,
			CreatedAt: timestamppb.Now(),
		}})
		break

	}
	return nil
}
