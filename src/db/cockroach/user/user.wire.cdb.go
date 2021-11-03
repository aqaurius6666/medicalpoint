//go:build wireinject
// +build wireinject
package user

import (
	"context"

	"github.com/google/wire"
	"github.com/medicalpoint/gateway/src/db/interface/user"
	"github.com/sirupsen/logrus"
	"github.com/sonntuet1997/medical-chain-utils/cockroach"
	"github.com/sonntuet1997/medical-chain-utils/common"
	"gorm.io/gorm"
)

func InitUserCDBRepo(ctx context.Context, logger *logrus.Logger, db *gorm.DB) (*UserCDBRepo, error) {
	wire.Build(
		wire.Value(cockroach.DBInterfaces{&user.User{}}),
		cockroach.InitCDBRepo,
		wire.Struct(new(UserCDBRepo), "CDBRepo"),
	)
	return &UserCDBRepo{}, nil
}

func InitUserCDBMockRepo(ctx context.Context, dsn string) (*UserCDBRepo, error) {
	wire.Build(
		cockroach.NewCDBConnection,
		InitUserCDBRepo,
		common.InitLoggerWithoutCLIContext,
	)
	return &UserCDBRepo{}, nil
}
