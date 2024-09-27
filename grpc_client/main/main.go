package main

import (
	"context"
	gen "github.com/golanshy/grpc-playground/gen"
	"google.golang.org/grpc"
	"log"
	"os"
	"strings"
	"time"
)

//goland:noinspection ALL
func main() {

	log.Print("Setting up GRPC Client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	defer conn.Close()

	client := gen.NewKeyValueClient(conn)

	var action, key, value string

	if len(os.Args) > 2 {
		action, key = os.Args[1], os.Args[2]
		value = strings.Join(os.Args[3:], " ")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call client.Get() or client.Put()

	switch action {

	case "get":
		r, err := client.Get(ctx, &gen.GetRequest{Key: key})
		if err != nil {
			log.Fatalf("Could not get value for key %s: %v", key, err)
		}
		log.Printf("Get %s returns %s", key, r.Value)

	case "put":
		_, err := client.Put(ctx, &gen.PutRequest{Key: key, Value: value})
		if err != nil {
			log.Fatalf("Could not get value for key %s: %v", key, err)
		}
		log.Printf("Put %s %s", key, value)

	default:
		log.Fatalf("Could not get value for key %s: not implemented", key)
	}
}
