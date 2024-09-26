package main

import (
	"context"
	gen "github.com/golanshy/grpc-playground/gen"
)

type server struct {
	gen.UnimplementedKeyValueServer
}

func (s *server) Get(ctx context.Context, r *gen.GetRequest) (*gen.GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented")
}

func (s *server) Put(ctx context.Context, r *gen.GetRequest) (*gen.GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented")
}

func (s *server) Delete(ctx context.Context, r *gen.GetRequest) (*gen.GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented")
}
