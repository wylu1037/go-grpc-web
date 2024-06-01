package ping

import (
	"go.uber.org/fx"
	"lattice-manager-grpc/app/module/ping/service"
)

var NewModule = fx.Options(
	fx.Provide(service.NewPingServiceServer),
)
