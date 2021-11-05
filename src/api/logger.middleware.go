package api

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type LoggerMiddleware struct {
	logger *logrus.Logger
}

func (l *LoggerMiddleware) Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// start time
		startTime := time.Now()
		// Processing request
		c.Next()

		// End time
		endTime := time.Now()
		// execution time
		latencyTime := endTime.Sub(startTime)
		// Request mode
		reqMethod := c.Request.Method
		// Request routing
		reqUri := c.Request.RequestURI
		// Status code
		statusCode := c.Writer.Status()
		// Request IP
		// clientIP := c.ClientIP()
		// Log format
		// l.logger.Infof("| %3d | %13v | %15s | %s | %s |",
		// 	statusCode,
		// 	latencyTime,
		// 	clientIP,
		// 	reqMethod,
		// 	reqUri,
		// )
		reqLogger := l.logger.WithFields(logrus.Fields{
			"status":  fmt.Sprintf("%3d", statusCode),
			"latency": fmt.Sprintf("%v", latencyTime),
			"method":  reqMethod,
			"path":    reqUri,
		})
		if body, ok := c.Get("body"); ok {
			reqLogger = reqLogger.WithField("body", string(body.([]byte)))
		}
		if err, ok := c.Get("error"); ok {
			reqLogger.Errorf("%+v", err)
		} else {
			reqLogger.Info()
		}
	}
}
