package middleware

import (
	"context"
	"github.com/samber/lo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// gRPC方法白名单
var methodWhiteList = []string{
	"/ping.PingService/Ping",
}

func Next(ctx context.Context) (bool, error) {
	if method, exist := grpc.Method(ctx); !exist {
		return false, status.Errorf(codes.NotFound, "method not found from grpc context")
	} else {
		return lo.IndexOf(methodWhiteList, method) != -1, nil
	}
}
