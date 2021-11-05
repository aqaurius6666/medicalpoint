//go:build wireinject
// +build wireinject
package api

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/medicalpoint/gateway/src/db"
	"github.com/medicalpoint/gateway/src/services/cosmos"
	"github.com/sirupsen/logrus"
)

type ApiServiceOptions struct {
	Logger         *logrus.Logger
	GateWayCDBRepo db.GateWayServiceRepo
	Chain          *cosmos.CosmosServiceClient
	Key            string
}

func InitApiService(ctx context.Context, opts ApiServiceOptions) (*GateWayServer, error) {
	wire.Build(
		wire.FieldsOf(&opts, "Logger", "GateWayCDBRepo", "Chain", "Key"),
		gin.New,
		wire.Struct(new(BlockchainApi), "*"),
		wire.Struct(new(BlockchainService), "*"),
		wire.Struct(new(LoggerMiddleware), "*"),
		wire.Struct(new(GateWayServer), "*"),
	)
	return &GateWayServer{}, nil
}
