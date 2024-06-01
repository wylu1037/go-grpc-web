package router

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"lattice-manager-grpc/gen/ping"
	tblockv1 "lattice-manager-grpc/gen/tblock/v1"
)

type Router struct {
	grpcServer          *grpc.Server
	pingServiceServer   ping.PingServiceServer
	tBlockServiceServer tblockv1.TBlockServiceServer
}

type handlerFunc func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error

var handlers []handlerFunc

func NewRouter(grpcServer *grpc.Server, pingServiceServer ping.PingServiceServer, tBlockServiceServer tblockv1.TBlockServiceServer) *Router {
	return &Router{
		grpcServer:          grpcServer,
		pingServiceServer:   pingServiceServer,
		tBlockServiceServer: tBlockServiceServer,
	}
}

func (r *Router) Register() {
	// gRPC call
	ping.RegisterPingServiceServer(r.grpcServer, r.pingServiceServer)
	tblockv1.RegisterTBlockServiceServer(r.grpcServer, r.tBlockServiceServer)

	// http call: gRPC-gateway
	handlers = append(handlers, ping.RegisterPingServiceHandler, tblockv1.RegisterTBlockServiceHandler)
}

func RegisterHandler(mux *runtime.ServeMux, grpcClientConn *grpc.ClientConn) {
	for _, h := range handlers {
		_ = h(context.Background(), mux, grpcClientConn)
	}
}
