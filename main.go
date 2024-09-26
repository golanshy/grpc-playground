package main

import (
	"context"
	pb "github.com/golanshy/grpc-playground/keyvalue"
)

type server struct {
	pb.UnimplementedKeyValueServer
}

func (s *server) Get(ctx context.Context, r *pb.GetRequest) (*pb.GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented")
}

func (s *server) Put(ctx context.Context, r *pb.GetRequest) (*pb.GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented")
}

func (s *server) Delete(ctx context.Context, r *pb.GetRequest) (*pb.GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented")
}
