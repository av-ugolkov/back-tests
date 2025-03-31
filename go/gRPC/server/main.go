package main

import (
	"context"
	"log"
	"net"

	pb "github.com/av-ugolkov/backend-examples/go/gRPC/proto"

	"google.golang.org/grpc"
)

//go:generate protoc --go_out=./proto --go-grpc_out=./proto --go_opt=paths= --go-grpc_opt=paths=source_relative ./proto/service.proto

func main() {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterTestServiceServer(srv, &server{})

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct {
	pb.UnimplementedTestServiceServer
}

func (s *server) SayHello(_ context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}
