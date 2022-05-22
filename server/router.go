/// router.go API 路由配置
package server

import (
	"freyja/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func newRouter() *gin.Engine {
	r := gin.New()

	r.Use(middleware.LoggerMiddleware())
	r.Use(gin.Recovery())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello world")
	})

	return r
}
