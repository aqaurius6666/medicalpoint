package api

import (
	"github.com/medicalpoint/gateway/src/db"

	"github.com/sirupsen/logrus"
)

type BlockchainService struct {
	db     db.GateWayServiceRepo
	logger *logrus.Logger
}
