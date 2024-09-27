package main

import (
	gen "github.com/golanshy/grpc-playground/gen"
	grpcServer "github.com/golanshy/grpc-playground/grpc_server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	log.Print("Setting up GRPC Server")

	s := grpc.NewServer()
	gen.RegisterKeyValueServer(s, &grpcServer.Server{})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
