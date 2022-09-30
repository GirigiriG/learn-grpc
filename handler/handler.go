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

func (p *Server) CreateProfileStream(x *profile.CreateRequest, stream profile.ProfileService_CreateProfileStreamServer) error {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 0)
		stream.Send(&profile.CreateResponse{Message: strconv.Itoa(i)})
	}
	return nil
}
