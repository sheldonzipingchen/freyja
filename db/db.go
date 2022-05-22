/// 数据库连接设置
package db

import (
	"fmt"
	"freyja/config"
	"freyja/lg"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	log := lg.GetLog()

	c := config.GetConfig()

	ip := c.GetString("db.ip")
	port := c.GetString("db.port")
	username := c.GetString("db.username")
	password := c.GetString("db.password")
	dbName := c.GetString("db.database")

	var dsn string
	if password == "" {
		dsn = fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", ip, username, dbName, port)

	} else {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", ip, username, password, dbName, port)

	}

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.WithFields(logrus.Fields{
			"ip":       ip,
			"port":     port,
			"username": username,
			"database": dbName,
			"error":    err,
		}).Fatal("connect to database fatal")

	}

}

/// GetDB 获取数据库
func GetDB() *gorm.DB {
	return db
}
