package bootstrap

import (
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io/fs"
	"lattice-manager-grpc/app/middleware"
	"lattice-manager-grpc/app/router"
	"lattice-manager-grpc/config"
	"lattice-manager-grpc/third_party"
	"mime"
	"net"
	"net/http"
	"strings"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func NewGRPCServer(config *config.Config) *grpc.Server {
	return grpc.NewServer(middleware.RegisterInterceptors(config))
}

func newGRPCClientConn() *grpc.ClientConn {
	conn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf(":%d", *port),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to dial grpc-gateway server")
	}
	return conn
}

func startGRPCGatewayServer() error {
	grpcClientConn := newGRPCClientConn()
	mux := runtime.NewServeMux()
	router.RegisterHandler(mux, grpcClientConn)
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", *port+1),
		Handler: http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			if strings.HasPrefix(request.URL.Path, "/api") {
				mux.ServeHTTP(writer, request)
				return
			}
			newSwaggerUIHandler().ServeHTTP(writer, request)
		}),
	}
	return server.ListenAndServe()
}

// new swagger-ui handler
// Returns:
//   - http.Handler: impl func ServeHTTP(ResponseWriter, *Request)
func newSwaggerUIHandler() http.Handler {
	_ = mime.AddExtensionType(".svg", "image/svg+xml")
	// Use subdirectory in embedded files
	subFS, err := fs.Sub(third_party.SwaggerUI, "swagger-ui")
	if err != nil {
		panic("swagger-ui couldn't create sub filesystem: " + err.Error())
	}
	return http.FileServer(http.FS(subFS))
}

// Start 启动 GRPC 服务
func Start(lifecycle fx.Lifecycle, grpcServer *grpc.Server, r *router.Router) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			flag.Parse()
			lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
			if err != nil {
				log.Fatal().Err(err).Msgf("failed to listen")
			}

			r.Register()

			go func() {
				if err := grpcServer.Serve(lis); err != nil {
					log.Fatal().Err(err).Msg("failed to serve")
				}
			}()

			// start grpc-gateway
			go func() {
				if err := startGRPCGatewayServer(); err != nil {
					log.Fatal().Msgf("failed to serve: %v", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
