/// log_middleware Gin日志中间件
package middleware

import (
	"freyja/lg"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

/// LoggerMiddleware Gin日志中间件
func LoggerMiddleware() gin.HandlerFunc {
	log := lg.GetLog()

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUrl := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		log.WithFields(logrus.Fields{
			"Status Code":    statusCode,
			"Latency Time":   latencyTime,
			"Client IP":      clientIP,
			"Request Method": reqMethod,
			"Request URL":    reqUrl,
		}).Info("Request Information.")

	}
}
