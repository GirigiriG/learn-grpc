package handler

import (
	"fmt"
	"io"
	"log"
	"logger/proto/logger"
)

type Server struct {
	logger.UnimplementedLoggerServiceServer
}

func (s *Server) Streamer(stream logger.LoggerService_StreamerServer) error {
	log.Println("Processing client data stream")

	res := ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&logger.Response{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v", err)
		}

		res += fmt.Sprintf("%s\n", req.Level)
	}
	return nil
}
