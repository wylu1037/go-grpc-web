package middleware

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"runtime"
)

func NewRecoveryInterceptor() grpc.UnaryServerInterceptor {
	// Define recoveryFunc to handle panic
	recoveryFunc := func(p any) (err error) {
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}
	// Shared options for the logger, with a custom gRPC code to log level function.
	opts := []RecoveryOption{
		WithRecoveryHandler(recoveryFunc),
	}

	return UnaryServerRecoveryInterceptor(opts...)
}

// RecoveryHandlerFunc is a function that recovers from the panic `p` by returning an `error`.
type RecoveryHandlerFunc func(p any) (err error)

// RecoveryHandlerFuncContext is a function that recovers from the panic `p` by returning an `error`.
// The context can be used to extract request scoped metadata and context values.
type RecoveryHandlerFuncContext func(ctx context.Context, p any) (err error)

// UnaryServerRecoveryInterceptor returns a new unary server interceptor for panic recovery.
func UnaryServerRecoveryInterceptor(opts ...RecoveryOption) grpc.UnaryServerInterceptor {
	o := evaluateOptions(opts)
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ any, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = recoverFrom(ctx, r, o.recoveryHandlerFunc)
			}
		}()

		return handler(ctx, req)
	}
}

func recoverFrom(ctx context.Context, p any, r RecoveryHandlerFuncContext) error {
	if r != nil {
		return r(ctx, p)
	}
	stack := make([]byte, 64<<10)
	stack = stack[:runtime.Stack(stack, false)]
	return &PanicError{Panic: p, Stack: stack}
}

type PanicError struct {
	Panic any
	Stack []byte
}

func (e *PanicError) Error() string {
	return fmt.Sprintf("panic caught: %v\n\n%s", e.Panic, e.Stack)
}

// options
var (
	defaultOptions = &options{
		recoveryHandlerFunc: nil,
	}
)

type options struct {
	recoveryHandlerFunc RecoveryHandlerFuncContext
}

func evaluateOptions(opts []RecoveryOption) *options {
	optCopy := &options{}
	*optCopy = *defaultOptions
	for _, o := range opts {
		o(optCopy)
	}
	return optCopy
}

type RecoveryOption func(*options)

// WithRecoveryHandler customizes the function for recovering from a panic.
func WithRecoveryHandler(f RecoveryHandlerFunc) RecoveryOption {
	return func(o *options) {
		o.recoveryHandlerFunc = RecoveryHandlerFuncContext(func(ctx context.Context, p any) error {
			return f(p)
		})
	}
}

// WithRecoveryHandlerContext customizes the function for recovering from a panic.
func WithRecoveryHandlerContext(f RecoveryHandlerFuncContext) RecoveryOption {
	return func(o *options) {
		o.recoveryHandlerFunc = f
	}
}
