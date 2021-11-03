//go:build wireinject
// +build wireinject

package api

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/medicalpoint/gateway/src/db"
	"github.com/sirupsen/logrus"
)

type ApiServiceOptions struct {
	Logger         *logrus.Logger
	GateWayCDBRepo db.GateWayServiceRepo
}

func InitApiService(ctx context.Context, opts ApiServiceOptions) (*GateWayServer, error) {
	wire.Build(
		wire.FieldsOf(&opts, "Logger", "GateWayCDBRepo"),
		gin.New,
		wire.Struct(new(BlockchainApi), "*"),
		wire.Struct(new(BlockchainService), "*"),
		wire.Struct(new(GateWayServer), "*"),
	)
	return &GateWayServer{}, nil
}
