package main

import (
	"context"
	"fmt"

	"buf.build/go/protovalidate"
	"github.com/gilperopiola/go-globant-grpc/pbs"
)

type service struct {
	pbs.UnimplementedAuthServiceServer
	pbs.UnimplementedUsersServiceServer

	validator protovalidate.Validator
}

func (s *service) Login(ctx context.Context, req *pbs.LoginRequest) (*pbs.LoginResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, err
	}
	return &pbs.LoginResponse{Token: "fake-token"}, nil
}

func (s *service) GetUsers(ctx context.Context, req *pbs.GetUsersRequest) (*pbs.GetUsersResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, err
	}
	usernames := make([]string, len(req.Ids))
	for i, id := range req.Ids {
		usernames[i] = fmt.Sprintf("user-%d", id)
	}
	return &pbs.GetUsersResponse{Usernames: usernames}, nil
}
