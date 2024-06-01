package main

import (
	"go.uber.org/fx"
	"lattice-manager-grpc/app/module/ping"
	"lattice-manager-grpc/app/module/tblock"
	"lattice-manager-grpc/app/router"
	"lattice-manager-grpc/config"
	"lattice-manager-grpc/internal/bootstrap"
)

func main() {
	fx.New(
		fx.Provide(config.NewConfig),
		fx.Provide(bootstrap.NewGRPCServer),
		fx.Provide(router.NewRouter),
		ping.NewModule,
		tblock.NewModule,
		fx.Invoke(bootstrap.Start),
	).Run()
}
