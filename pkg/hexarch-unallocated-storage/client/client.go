package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	pb "github.com/av-ugolkov/backend-examples/pkg/unallocated-storage/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithIdleTimeout(time.Second))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewKeyValueClient(conn)
	var action, key, value string

	if len(os.Args) > 2 {
		action, key = os.Args[1], os.Args[2]
		value = strings.Join(os.Args[3:], " ")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	switch action {
	case "get":
		r, err := client.Get(ctx, &pb.GetRequest{Key: key})
		if err != nil {
			log.Fatalf("could not get value for key %s: %v\n", key, err)
		}
		log.Printf("Get %s returns: %s", key, r.GetValue())
	case "put":
		_, err := client.Put(ctx, &pb.PutRequest{Key: key, Value: value})
		if err != nil {
			log.Fatalf("could not put key %s: %v\n", key, err)
		}
		log.Printf("Put %s", key)
	default:
		log.Fatalf("Syntax: go run [get|put] KEY VALUE...")
	}
}
