package bootstrap

import (
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"lattice-manager-grpc/app/router"
	"log"
	"net"
	"net/http"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func NewGRPCServer() *grpc.Server {
	return grpc.NewServer()
}

func newGRPCClientConn() *grpc.ClientConn {
	conn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf(":%d", *port),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial grpc-gateway server:", err)
	}
	return conn
}

func startGRPCGatewayServer() error {
	grpcClientConn := newGRPCClientConn()
	mux := runtime.NewServeMux()
	router.RegisterHandler(mux, grpcClientConn)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", *port+1),
		Handler: mux,
	}
	return server.ListenAndServe()
}

// Start 启动 GRPC 服务
func Start(lifecycle fx.Lifecycle, grpcServer *grpc.Server, r *router.Router) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			flag.Parse()
			lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}

			r.Register()

			go func() {
				if err := grpcServer.Serve(lis); err != nil {
					log.Fatalf("failed to serve: %v", err)
				}
			}()

			// start grpc-gateway
			go func() {
				if err := startGRPCGatewayServer(); err != nil {
					log.Fatalf("failed to serve: %v", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
