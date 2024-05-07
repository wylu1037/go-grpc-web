package middleware

import (
	"context"
	"fmt"
	"github.com/juju/ratelimit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewBucket() *ratelimit.Bucket {
	return ratelimit.NewBucketWithRate(2000, 20000)
}

func NewRateLimitInterceptor() grpc.UnaryServerInterceptor {
	limit := &limiter{
		bucket: NewBucket(),
	}

	return UnaryServerRateLimitInterceptor(limit)
}

// Limiter defines the interface to perform request rate limiting.
// If Limit function returns an error, the request will be rejected with the gRPC codes.ResourceExhausted and the provided error.
// Otherwise, the request will pass.
type Limiter interface {
	Limit(ctx context.Context) error
}

type limiter struct {
	bucket *ratelimit.Bucket
}

func (l *limiter) Limit(ctx context.Context) error {
	n := l.bucket.TakeAvailable(1)
	if n == 0 {
		return fmt.Errorf("reached Rate-Limiting %d", l.bucket.Available())
	}
	// Rate limit isn't reached.
	return nil
}

// UnaryServerRateLimitInterceptor returns a new unary server interceptors that performs request rate limiting.
func UnaryServerRateLimitInterceptor(limiter Limiter) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		if err := limiter.Limit(ctx); err != nil {
			return nil, status.Errorf(codes.ResourceExhausted, "%s is rejected by grpc_ratelimit middleware, please retry later. %s", info.FullMethod, err)
		}
		return handler(ctx, req)
	}
}

// StreamServerInterceptor returns a new stream server interceptor that performs rate limiting on the request.
func StreamServerInterceptor(limiter Limiter) grpc.StreamServerInterceptor {
	return func(srv any, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if err := limiter.Limit(stream.Context()); err != nil {
			return status.Errorf(codes.ResourceExhausted, "%s is rejected by grpc_ratelimit middleware, please retry later. %s", info.FullMethod, err)
		}
		return handler(srv, stream)
	}
}

// UnaryClientInterceptor returns a new unary client interceptor that performs rate limiting on the request on the
// client side.
// This can be helpful for clients that want to limit the number of requests they send in a given time, potentially
// saving cost.
func UnaryClientInterceptor(limiter Limiter) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if err := limiter.Limit(ctx); err != nil {
			return status.Errorf(codes.ResourceExhausted, "%s is rejected by grpc_ratelimit middleware, please retry later. %s", method, err)
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

// StreamClientInterceptor returns a new stream client interceptor that performs rate limiting on the request on the
// client side.
// This can be helpful for clients that want to limit the number of requests they send in a given time, potentially
// saving cost.
func StreamClientInterceptor(limiter Limiter) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		if err := limiter.Limit(ctx); err != nil {
			return nil, status.Errorf(codes.ResourceExhausted, "%s is rejected by grpc_ratelimit middleware, please retry later. %s", method, err)
		}
		return streamer(ctx, desc, cc, method, opts...)
	}
}
