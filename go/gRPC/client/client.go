package main

import (
	"context"
	"log"

	pb "github.com/av-ugolkov/backend-examples/go/gRPC/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	*grpc.ClientConn

	pb.TestServiceClient
}

func NewClient() *Client {
	conn, err := grpc.NewClient(":5000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := pb.NewTestServiceClient(conn)
	return &Client{conn, c}
}

func (c *Client) Close() {
	c.ClientConn.Close()
}

func (c *Client) SayHello(ctx context.Context, name string) (string, error) {
	r, err := c.TestServiceClient.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		return "", err
	}

	return r.GetMessage(), nil
}
