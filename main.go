package main

import (
	pb "github.com/golanshy/grpc-playground/gen"
	"github.com/golanshy/grpc-playground/pb_server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	pb.RegisterKeyValueServer(s, &pb_server.Server{})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
