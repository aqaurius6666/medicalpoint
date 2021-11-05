//go:build wireinject
// +build wireinject
package main

import (
	"context"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/google/wire"
	"github.com/medicalpoint/gateway/src/api"
	"github.com/medicalpoint/gateway/src/db"
	"github.com/medicalpoint/gateway/src/services/cosmos"
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
	Key        string
}

func InitGateWayServer(ctx context.Context, opts AppOptions) (*api.GateWayServer, error) {
	wire.Build(
		wire.FieldsOf(&opts, "dbDsn", "cliContext", "KeyRing", "CosmosEp", "Mne", "ChainId", "Key"),
		db.InitGateWayServiceRepo,
		common.InitLogger,
		api.InitApiService,
		cosmos.InitCosmosServiceClient,
		wire.Struct(new(api.ApiServiceOptions), "*"),
	)

	return &api.GateWayServer{}, nil
}
