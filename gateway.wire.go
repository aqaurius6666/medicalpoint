//go:build wireinject
// +build wireinject

package main

import (
	"api"
	"context"

	"github.com/sonntuet1997/medical-chain-utils/common"
	"github.com/urfave/cli/v2"
)

type AppOptions struct {
	dbDsn      db.DBDsn
	cliContext *cli.Context
	ChainId    cosmos.ChainID
	KeyRing    keyring.UnsafeKeyring
	CosmosEp   cosmos.CosmosEndpoint
	Mne        cosmos.Mnemonic
}

func InitGateWayServer(ctx context.Context, opts AppOptions) (*api.GateWayServer, error) {
	wire.Build(
		wire.FieldsOf(&opts, "authServiceDsn", "dbDsn", "cliContext", "AllowKill", "FBAccountPath", "FilePath", "ChainId", "KeyRing", "CosmosEp", "Mne", "ProjectId"),
		auth.NewAuthServiceClient,
		db.InitMainServiceRepo,
		common.InitLogger,
		api.InitApiService,
		fbcm.InitFBService,
		cosmos.InitCosmosServiceClient,
		wire.Struct(new(api.ApiServiceOptions), "*"),
	)

	return &api.MainServer{}, nil
}
