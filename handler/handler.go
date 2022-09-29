package handler

import (
	"context"
	"fmt"
	"profile/proto/profile"
)

type Profile struct {
	profile.UnimplementedProfileServiceServer
}

func (p *Profile) Create(ctx context.Context, req *profile.CreateRequest) (*profile.CreateResponse, error) {
	return &profile.CreateResponse{Message: fmt.Sprintf("created profile for %s", req.GetName())}, nil
}

func CreateProfileStream(ctx context.Context, in *profile.CreateRequest) (profile.ProfileService_CreateProfileStreamClient, error) {
	for i := 0; i < 10; i++ {
		//TODO: Figure out why stream.Send is not being imported
		stream.Send("")
	}
	return nil, nil
}
