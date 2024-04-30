package router

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"lattice-manager-grpc/gen/helloworld"
	tblockv1 "lattice-manager-grpc/gen/tblock/v1"
)

type Router struct {
	grpcServer *grpc.Server
	chat       helloworld.GreeterServer
}

type handlerFunc func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error

var handlers []handlerFunc

func NewRouter(grpcServer *grpc.Server, chat helloworld.GreeterServer) *Router {
	return &Router{
		grpcServer: grpcServer,
		chat:       chat,
	}
}

func (r *Router) Register() {
	helloworld.RegisterGreeterServer(r.grpcServer, r.chat)

	handlers = append(handlers, helloworld.RegisterGreeterHandler, tblockv1.RegisterTBlockServiceHandler)
}

func RegisterHandler(mux *runtime.ServeMux, grpcClientConn *grpc.ClientConn) {
	for _, h := range handlers {
		_ = h(context.Background(), mux, grpcClientConn)
	}
}
