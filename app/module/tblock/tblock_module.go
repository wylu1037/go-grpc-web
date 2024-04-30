package tblock

import (
	"go.uber.org/fx"
	"lattice-manager-grpc/app/module/tblock/service"
)

var NewModule = fx.Options(
	fx.Provide(service.NewTBlockServer),
)
