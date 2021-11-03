package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type GateWayServer struct {
	G          *gin.Engine
	blockchain *BlockchainApi
}

func (s *GateWayServer) RegisterEndpoint() {
	s.G.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Authorization", "Content-Type", "User-Agent"},
	}))

	// s.G.Use(gin.RecoveryWithWriter(s.loggerMiddleware.logger.Out))
	s.G.Use(gin.Recovery())
	// index := s.G.Group("/")

}
