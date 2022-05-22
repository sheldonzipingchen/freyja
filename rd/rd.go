/// rd.go redis客户端配置
package rd

import (
	"fmt"
	"freyja/config"
	"freyja/lg"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

var rdb *redis.Client

/// Init 初始化Redis客户端的连接
func Init() {
	c := config.GetConfig()
	log := lg.GetLog()

	ip := c.GetString("redis.ip")
	if ip == "" {
		ip = "127.0.0.1"
	}

	port := c.GetString("redis.port")
	if port == "" {
		port = "6379"
	}

	addr := fmt.Sprintf("%s:%s", ip, port)

	password := c.GetString("redis.password")
	dbIndex := c.GetInt("redis.db")

	log.WithFields(logrus.Fields{
		"Redis Server IP":       ip,
		"Redis Server Port":     port,
		"Redis Server DB Index": dbIndex,
	}).Info("Connect Redis Server Parameters.")

	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       dbIndex,
	})

}

/// GetRDB 返回Redis客户端连接
func GetRDB() *redis.Client {
	return rdb
}
