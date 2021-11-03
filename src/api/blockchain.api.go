package api

import "github.com/sirupsen/logrus"

type BlockchainApi struct {
	blockchainService *BlockchainService

	logger *logrus.Logger
}
