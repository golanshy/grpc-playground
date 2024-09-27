package pb_server

import (
	"context"
	"errors"
	gen "github.com/golanshy/grpc-playground/gen"
)

type Server struct {
	gen.UnimplementedKeyValueServer
}

func (s *Server) Get(ctx context.Context, r *gen.GetRequest) (*gen.GetResponse, error) {
	return nil, errors.New("method not implemented")
}

func (s *Server) Put(ctx context.Context, r *gen.PutRequest) (*gen.PutResponse, error) {
	return nil, errors.New("method not implemented")
}

func (s *Server) Delete(ctx context.Context, r *gen.DeleteRequest) (*gen.DeleteResponse, error) {
	return nil, errors.New("method not implemented")
}
