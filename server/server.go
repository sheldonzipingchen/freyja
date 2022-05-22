/// server.go API 服务初始化
package server

import (
	"freyja/config"
	"freyja/lg"
	"net/http"

	"github.com/sirupsen/logrus"
)

func Init() {
	c := config.GetConfig()
	log := lg.GetLog()

	r := newRouter()

	port := c.GetString("server.port")
	if port == "" {
		// 设置默认 web api 端口为 8000
		port = "8000"
	}

	log.WithFields(logrus.Fields{
		"port": port,
	}).Info("HTTP API Server is listening...")

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("fataled to start http server.")
	}
}
