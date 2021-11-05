//go:build wireinject
// +build wireinject

package cosmos

import (
	"context"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"github.com/sonntuet1997/medical-chain-utils/common"
)

func InitCosmosServiceClient(ctx context.Context, logger *logrus.Logger, ChainId ChainID, KeyRing keyring.UnsafeKeyring, CosmosEp CosmosEndpoint, Mne Mnemonic) (*CosmosServiceClient, error) {
	wire.Build(
		wire.Struct(new(CosmosOpts), "*"),
		NewCosmosServiceClient,
	)
	return &CosmosServiceClient{}, nil
}

func InitTestCosmosServiceClient(ctx context.Context, ChainId ChainID, KeyRing keyring.UnsafeKeyring, CosmosEp CosmosEndpoint, Mne Mnemonic) (*CosmosServiceClient, error) {

	wire.Build(
		wire.Struct(new(CosmosOpts), "*"),
		NewCosmosServiceClient,
		common.InitLoggerWithoutCLIContext,
	)
	return &CosmosServiceClient{}, nil
}
