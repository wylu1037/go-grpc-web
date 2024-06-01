package service

import (
	"context"
	"lattice-manager-grpc/gen/ping"
)

func NewPingServiceServer() ping.PingServiceServer {
	return &pingServiceServer{}
}

type pingServiceServer struct {
	ping.UnimplementedPingServiceServer
}

func (svc *pingServiceServer) Ping(context.Context, *ping.PingRequest) (*ping.PongResponse, error) {
	return &ping.PongResponse{
		Reply: "🎉🎉🎉pong!",
	}, nil
}
