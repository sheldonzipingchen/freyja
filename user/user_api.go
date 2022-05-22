/// user_api.go API定义
package user

import (
	"context"
	"fmt"
	"freyja/lg"
	"freyja/web"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

/// UserAPI API定义
type UserAPI struct {
	DB  *gorm.DB
	RDB *redis.Client
}

/// FindUser 查找用户信息(登录API方法)
func (api *UserAPI) FindUser(c *gin.Context) {
	resp := &web.Response{}

	log := lg.GetLog()

	var findUserRequest findUserRequest

	if err := c.ShouldBind(&findUserRequest); err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Error("FindUser Request Error.")

		resp.Success = false
		resp.Message = "用户登录错误"
		c.JSON(http.StatusOK, resp)
		return
	}

	paramLog := log.WithFields(logrus.Fields{
		"findUserRequest": findUserRequest,
	})

	var localAuth LocalAuth
	err := api.DB.Where("email = ?", findUserRequest.Email).First(&localAuth).Error
	if err != nil {
		paramLog.WithFields(logrus.Fields{
			"error": err,
		}).Warn("User is not exists.")
		resp.Success = false
		resp.Message = "用户名或密码错误"
		c.JSON(http.StatusOK, resp)
		return
	}

	if !localAuth.CheckPasswordHash(findUserRequest.Password) {
		paramLog.WithFields(logrus.Fields{
			"error": err,
		}).Warn("Password is not correct.")
		resp.Success = false
		resp.Message = "用户名或密码错误"
		c.JSON(http.StatusOK, resp)
		return
	}

	u4 := uuid.New()
	accessToken := u4.String()

	key := fmt.Sprintf("freyja-user-%s", accessToken)
	ctx := context.Background()
	_, err = api.RDB.SetNX(ctx, key, localAuth.UserID, 24*time.Hour).Result()
	if err != nil {
		paramLog.WithFields(logrus.Fields{
			"error": err,
		}).Warn("Save accesstoken to redis server error.")
		resp.Success = false
		resp.Message = "服务出错，请联系管理员."
		c.JSON(http.StatusOK, resp)

		return
	}

	loginResponseEntity := LoginResponseEntity{
		AccessToken: accessToken,
	}
	resp.Success = true
	resp.Message = "用户登录成功"
	resp.Data = loginResponseEntity

	c.JSON(http.StatusOK, resp)
}
