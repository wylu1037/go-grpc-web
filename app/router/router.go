package router

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"lattice-manager-grpc/gen/helloworld"
	tblockv1 "lattice-manager-grpc/gen/tblock/v1"
)

type Router struct {
	grpcServer    *grpc.Server
	greeterServer helloworld.GreeterServer
	tBlockServer  tblockv1.TBlockServiceServer
}

type handlerFunc func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error

var handlers []handlerFunc

func NewRouter(grpcServer *grpc.Server, greeterServer helloworld.GreeterServer, tBlockServer tblockv1.TBlockServiceServer) *Router {
	return &Router{
		grpcServer:    grpcServer,
		greeterServer: greeterServer,
		tBlockServer:  tBlockServer,
	}
}

func (r *Router) Register() {
	// gRPC call
	helloworld.RegisterGreeterServer(r.grpcServer, r.greeterServer)
	tblockv1.RegisterTBlockServiceServer(r.grpcServer, r.tBlockServer)

	// http call: gRPC-gateway
	handlers = append(handlers, helloworld.RegisterGreeterHandler, tblockv1.RegisterTBlockServiceHandler)
}

func RegisterHandler(mux *runtime.ServeMux, grpcClientConn *grpc.ClientConn) {
	for _, h := range handlers {
		_ = h(context.Background(), mux, grpcClientConn)
	}
}
