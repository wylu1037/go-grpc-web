package middleware

import (
	"context"
	"github.com/bufbuild/protovalidate-go"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// NewValidatorInterceptor
// See https://github.com/grpc-ecosystem/go-grpc-middleware/blob/main/interceptors/validator/doc.go
// See https://github.com/bufbuild/protovalidate
// See https://github.com/bufbuild/protovalidate-go
func NewValidatorInterceptor() grpc.UnaryServerInterceptor {
	logErr := func(ctx context.Context, err error) {
		log.Error().Err(err).Msgf("middleware: failed to validate")
	}
	goValidator, err := protovalidate.New()
	if err != nil {
		log.Error().Err(err).Msgf("middleware: failed to new protovalidate")
	}
	return UnaryServerInterceptor(WithProtoValidate(goValidator), WithOnValidationErrCallback(logErr))
}

// UnaryServerInterceptor returns a new unary server interceptor that validates incoming messages.
//
// Invalid messages will be rejected with `InvalidArgument` before reaching any userspace handlers.
// Note that generated codes prior to protoc-gen-validate v0.6.0 do not provide an all-validation
// interface. In this case the interceptor fallbacks to legacy validation and `all` is ignored.
func UnaryServerInterceptor(opts ...Option) grpc.UnaryServerInterceptor {
	o := evaluateOpts(opts)
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if err := validate(ctx, req, o.protoValidate, o.onValidationErrCallback); err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

func validate(ctx context.Context, reqOrRes interface{}, protoValidate ProtoValidateFunc, onValidationErrCallback OnValidationErrCallback) (err error) {
	message, ok := reqOrRes.(proto.Message)
	if !ok {
		return nil
	}
	err = protoValidate(message)

	if err == nil {
		return nil
	}

	if onValidationErrCallback != nil {
		onValidationErrCallback(ctx, err)
	}
	return status.Error(codes.InvalidArgument, err.Error())
}

type options struct {
	protoValidate           ProtoValidateFunc
	onValidationErrCallback OnValidationErrCallback
}
type Option func(*options)

func evaluateOpts(opts []Option) *options {
	optCopy := &options{}
	for _, o := range opts {
		o(optCopy)
	}
	return optCopy
}

type OnValidationErrCallback func(ctx context.Context, err error)
type ProtoValidateFunc func(message proto.Message) error

// WithOnValidationErrCallback registers function that will be invoked on validation error(s).
func WithOnValidationErrCallback(onValidationErrCallback OnValidationErrCallback) Option {
	return func(o *options) {
		o.onValidationErrCallback = onValidationErrCallback
	}
}

// WithProtoValidate validate proto
func WithProtoValidate(v *protovalidate.Validator) Option {
	return func(o *options) {
		o.protoValidate = v.Validate
	}
}
