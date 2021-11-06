package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/medicalpoint/gateway/src/db"
	"github.com/medicalpoint/gateway/src/services/cosmos"
)

type GateWayServer struct {
	G          *gin.Engine
	blockchain *BlockchainApi
	Repo       db.GateWayServiceRepo
	chain      *cosmos.CosmosServiceClient
	logger     *LoggerMiddleware
}

func (s *GateWayServer) RegisterEndpoint() {
	s.G.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Authorization", "Content-Type", "User-Agent"},
	}))
	s.G.Use(s.logger.Logger())
	s.G.Use(gin.Recovery())
	medicalpoint := s.G.Group("/medicalpoint")
	medicalpoint.GET("", s.blockchain.HandleAdminGet)
	medicalpoint.POST("/users", s.blockchain.HandleUserPost)
	medicalpoint.GET("/balance", s.blockchain.HandleBalanceGet)
	medicalpoint.POST("/super-admin", s.blockchain.HandleSuperAdminPost)
	medicalpoint.POST("/mint", s.blockchain.HandleMintPost)
	medicalpoint.POST("/burn", s.blockchain.HandleBurnPost)
	medicalpoint.POST("/transfer", s.blockchain.HandleTransferPost)
	medicalpoint.POST("/admin-transfer", s.blockchain.HandleAdminTransferPost)
	medicalpoint.POST("/admin", s.blockchain.HandleAdminPost)
	medicalpoint.DELETE("/admin", s.blockchain.HandleAdminDelete)
	medicalpoint.GET("/system-balance", s.blockchain.HandleAdminDelete)
}
