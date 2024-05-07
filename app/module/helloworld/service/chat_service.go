package service

import (
	"context"
	"github.com/rs/zerolog/log"
	"lattice-manager-grpc/gen/helloworld"
)

func NewHelloWorldServer() helloworld.GreeterServer {
	return &chatServer{}
}

type chatServer struct {
	helloworld.UnimplementedGreeterServer
}

func (s *chatServer) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Info().Msgf("Received: %v", in.GetName())
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}
