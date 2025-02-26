package main

import (
	"context"
	"log"
	"net"

	"github.com/av-ugolkov/backend-examples/pkg/unallocated-storage/core"
	pb "github.com/av-ugolkov/backend-examples/pkg/unallocated-storage/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedKeyValueServer
}

func (s *server) Get(ctx context.Context, r *pb.GetRequest) (*pb.GetResponse, error) {
	log.Printf("Received GET key=%v", r.GetKey())
	value, err := core.Get(r.GetKey())
	return &pb.GetResponse{Value: value}, err
}

func (s *server) Put(ctx context.Context, r *pb.PutRequest) (*pb.PutResponse, error) {
	log.Printf("Received PUT key=%v", r.GetKey())
	err := core.Put(r.GetKey(), r.GetValue())
	return &pb.PutResponse{}, err
}

func main() {
	s := grpc.NewServer()
	pb.RegisterKeyValueServer(s, &server{})

	listn, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := s.Serve(listn); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
