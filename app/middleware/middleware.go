package middleware

import (
	"google.golang.org/grpc"
	"lattice-manager-grpc/config"
)

func RegisterInterceptors(config *config.Config) grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(NewRecoveryInterceptor(), NewRateLimitInterceptor(), NewLoggingInterceptor(config))
}
