package middleware

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/ratelimit"
	rl "github.com/juju/ratelimit"
	"google.golang.org/grpc"
)

func NewBucket() *rl.Bucket {
	return rl.NewBucketWithRate(100, 1000)
}

func NewRateLimitInterceptor() grpc.UnaryServerInterceptor {
	limit := &limiter{
		bucket: NewBucket(),
	}

	return ratelimit.UnaryServerInterceptor(limit)
}

type limiter struct {
	bucket *rl.Bucket
}

func (l *limiter) Limit(_ context.Context) error {
	n := l.bucket.TakeAvailable(1)
	if n == 0 {
		return fmt.Errorf("reached Rate-Limiting %d", l.bucket.Available())
	}
	// Rate limit isn't reached.
	return nil
}
