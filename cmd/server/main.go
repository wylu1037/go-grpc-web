package main

import (
	"go.uber.org/fx"
	"lattice-manager-grpc/app/module/helloworld"
	"lattice-manager-grpc/app/module/tblock"
	"lattice-manager-grpc/app/router"
	"lattice-manager-grpc/internal/bootstrap"
)

func main() {
	fx.New(
		fx.Provide(bootstrap.NewGRPCServer),
		fx.Provide(router.NewRouter),
		helloworld.NewModule,
		tblock.NewModule,
		fx.Invoke(bootstrap.Start),
	).Run()
}
