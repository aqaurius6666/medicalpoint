package db

import (
	"context"
	"net/url"

	"github.com/medicalpoint/gateway/src/db/cockroach"
	"github.com/medicalpoint/gateway/src/db/interface/user"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

type GateWayServiceRepo interface {
	user.UserRepo
}

type DBDsn string

func InitGateWayServiceRepo(ctx context.Context, logger *logrus.Logger, dsn DBDsn) (GateWayServiceRepo, error) {
	uri, err := url.Parse(string(dsn))
	if err != nil {
		return nil, xerrors.Errorf("could not parse DB URI: %w", err)
	}

	switch uri.Scheme {
	case "in-memory":
		logger.Info("using in-memory graph")
		return nil, xerrors.Errorf("Not implemented!", err)
	case "postgresql":
		return cockroach.InitGatewayCDBRepo(ctx, cockroach.GatewayCDBRepoOptions{
			Dsn:    string(dsn),
			Logger: logger,
		})
	default:
		return nil, xerrors.Errorf("unsupported DB URI scheme: %q", uri.Scheme)
	}
}
