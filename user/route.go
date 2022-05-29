// route.go API 路由定义
package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// UserAPP 用户API路由设置
func UserAPP(db *gorm.DB, rdb *redis.Client, router *gin.RouterGroup) {
	db.AutoMigrate(&LocalAuth{})
	db.AutoMigrate(&User{})

	userAPI := &UserAPI{DB: db, RDB: rdb}

	router.POST("/users/login", userAPI.FindUser)
	router.POST("/users/register", userAPI.CreateUser)
}
