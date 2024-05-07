package middleware

import "google.golang.org/grpc"

func RegisterInterceptors() grpc.ServerOption {
	return NewRecoveryInterceptor()
}
