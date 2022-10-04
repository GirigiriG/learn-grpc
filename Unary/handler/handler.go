package handler

import (
	"context"
	"fmt"
	"profile/proto/profile"
	"strconv"
	"time"
)

type Server struct {
	profile.UnimplementedProfileServiceServer
}

func (p *Server) Create(ctx context.Context, req *profile.CreateRequest) (*profile.CreateResponse, error) {
	return &profile.CreateResponse{Message: fmt.Sprintf("created profile for %s", req.GetName())}, nil
}

// Send stream of data from server
func (p *Server) CreateProfileStream(req *profile.CreateRequest, stream profile.ProfileService_CreateProfileStreamServer) error {
	i := 0
	for {
		time.Sleep(time.Millisecond * 300)
		stream.Send(&profile.CreateResponse{Message: strconv.Itoa(i) + " " + req.Name})
		i++
	}
	return nil
}
