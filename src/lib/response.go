package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sonntuet1997/medical-chain-utils/common"
)

func ErrBadRequest(g *gin.Context, err error) {
	g.Set("error", err)
	err = common.Unwrap(err)
	g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}

func ErrInternalServerError(g *gin.Context, err error) {
	g.Set("error", err)
	err = common.Unwrap(err)
	g.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
}

func ErrUnauthorized(g *gin.Context, err error) {
	g.Set("error", err)
	err = common.Unwrap(err)
	g.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error": err.Error(),
	})
}

func Success(g *gin.Context, data interface{}) {
	g.AbortWithStatusJSON(http.StatusOK, data)
}
