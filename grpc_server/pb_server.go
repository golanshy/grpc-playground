package pb_server

import (
	"context"
	"errors"
	pb "github.com/golanshy/grpc-playground/gen"
)

type Server struct {
	pb.UnimplementedKeyValueServer
}

func (s *Server) Get(ctx context.Context, r *pb.GetRequest) (*pb.GetResponse, error) {
	return nil, errors.New("method not implemented")
}

func (s *Server) Put(ctx context.Context, r *pb.PutRequest) (*pb.PutResponse, error) {
	return nil, errors.New("method not implemented")
}

func (s *Server) Delete(ctx context.Context, r *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	return nil, errors.New("method not implemented")
}
