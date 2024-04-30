package service

import (
	"context"
	"lattice-manager-grpc/gen/helloworld"
	"log"
)

func NewHelloWorldServer() helloworld.GreeterServer {
	return &chatServer{}
}

type chatServer struct {
	helloworld.UnimplementedGreeterServer
}

func (s *chatServer) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}
