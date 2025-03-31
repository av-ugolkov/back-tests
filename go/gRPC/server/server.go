package main

import (
	"context"
	"log"
	"net"

	pb "github.com/av-ugolkov/backend-examples/go/gRPC/proto"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedTestServiceServer

	grpcSrv *grpc.Server
}

func New() *Server {
	grpcSrv := grpc.NewServer()
	srv := &Server{
		grpcSrv: grpcSrv,
	}
	pb.RegisterTestServiceServer(grpcSrv, srv)
	return srv
}

func (s *Server) Serve(lis net.Listener) error {
	if err := s.grpcSrv.Serve(lis); err != nil {
		return err
	}
	return nil
}

func (s *Server) GracefulStop() {
	s.grpcSrv.GracefulStop()
}

func (s *Server) SayHello(_ context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}
