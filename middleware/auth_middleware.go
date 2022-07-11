package middleware

import (
	"context"
	"freyja/lg"
	"freyja/rd"
	"freyja/web"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// AuthMiddleware Gin身份验证中间件
func AuthMiddleware() gin.HandlerFunc {
	log := lg.GetLog()
	redisclient := rd.GetRDB()

	return func(c *gin.Context) {
		tokenID := c.GetHeader("authorization")
		log.WithFields(logrus.Fields{
			"token": tokenID,
		}).Debug("Check Auth Information.")

		if tokenID == "" {
			c.Abort()
			resp := web.NewFatalResponse(nil, "")
			c.JSON(http.StatusForbidden, resp)
		} else {
			ctx := context.Background()
			_, err := redisclient.Get(ctx, tokenID).Result()
			if err != nil {
				c.Abort()
				resp := web.NewFatalResponse(nil, "用户未登录")
				c.JSON(http.StatusForbidden, resp)
			} else {
				c.Next()
			}
		}
	}
}
