package middleware

import (
	"google.golang.org/grpc"
	"lattice-manager-grpc/config"
)

var unaryServerInterceptors []grpc.UnaryServerInterceptor

func RegisterInterceptors(cfg *config.Config) grpc.ServerOption {
	register(NewRecoveryInterceptor(), NewLoggingInterceptor(cfg), NewValidatorInterceptor())

	if cfg.Middleware.Limiter.Enable {
		register(NewRateLimitInterceptor())
	}
	if cfg.Middleware.Jwt.Enable {
		register(NewAuthInterceptor())
	}

	return grpc.ChainUnaryInterceptor(unaryServerInterceptors...)
}

func register(unaryServerInterceptor ...grpc.UnaryServerInterceptor) {
	unaryServerInterceptors = append(unaryServerInterceptors, unaryServerInterceptor...)
}
