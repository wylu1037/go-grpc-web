package helloworld

import (
	"go.uber.org/fx"
	"lattice-manager-grpc/app/module/helloworld/service"
)

var NewModule = fx.Options(
	fx.Provide(service.NewHelloWorldServer),
)
