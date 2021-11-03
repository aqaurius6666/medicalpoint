//go:build wireinject
// +build wireinject

package cockroach

import (
	"context"

	"github.com/medicalpoint/gateway/src/db/cockroach/user"
	"github.com/sonntuet1997/medical-chain-utils/cockroach"

	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

type GatewayCDBRepoOptions struct {
	Dsn    string
	Logger *logrus.Logger
}

func InitGatewayCDBRepo(ctx context.Context, opts GatewayCDBRepoOptions) (*GatewayCDBRepo, error) {
	wire.Build(
		wire.FieldsOf(&opts, "Dsn", "Logger"),
		user.InitUserCDBRepo,
		cockroach.NewCDBConnection,
		wire.Struct(new(GatewayCDBRepo), "*"),
	)
	return &GatewayCDBRepo{}, nil
}

// func InitGatewayCDBRepo(ctx context.Context, opts GatewayCDBRepoOptions) (*GatewayCDBRepo, error) {
// 	GatewayCDBRepo, err := initGatewayCDBRepo(ctx, opts)
// 	if err != nil {
// 		return nil, err
// 	}
// 	GatewayCDBRepo.Interfaces = []interface{}{
// 		&iUser.User{}, &iContact.Contact{},
// 		&iDevice.Device{}, &iThirdService.ThirdService{},
// 		&iConnect.Connect{}, &iRequest.Request{},
// 		&iNotificationQueue.NotificationQueue{},
// 		&iLog.Log{}, &iAdmin.Admin{},
// 		&iProvider.Provider{}, iAdminLog.AdminLog{},
// 	}
// 	return GatewayCDBRepo, nil
// }
