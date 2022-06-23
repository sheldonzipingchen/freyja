// user_api.go API定义
package user

import (
	"context"
	"errors"
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

// UserAPI API定义
type UserAPI struct {
	DB  *gorm.DB
	RDB *redis.Client
}

// FindUser 查找用户信息(登录API方法)
func (api *UserAPI) FindUser(c *gin.Context) {
	resp := &web.Response{}

	log := lg.GetLog()

	var findUserRequest findUserRequest

	if err := c.ShouldBind(&findUserRequest); err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Error("FindUser Request Error.")

		resp.Success = false
		resp.Code = web.PARAM_ERR
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
		resp.Code = web.SYS_ERR
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
		resp.Code = web.SYS_ERR
		resp.Message = "服务出错，请联系管理员."
		c.JSON(http.StatusOK, resp)

		return
	}

	loginResponseEntity := LoginResponseEntity{
		AccessToken: accessToken,
	}
	resp.Success = true
	resp.Code = web.SUCCESS
	resp.Message = "用户登录成功"
	resp.Data = loginResponseEntity

	c.JSON(http.StatusOK, resp)
}

// CreateUser 注册用户
func (api *UserAPI) CreateUser(c *gin.Context) {
	log := lg.GetLog()
	resp := &web.Response{}

	var userCreateRequest userCreateRequest
	if err := c.ShouldBind(&userCreateRequest); err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Error("CreateUser Request Error.")

		resp.Success = false
		resp.Code = web.PARAM_ERR
		resp.Message = "请求参数有误，请检查后再重新提交"
		c.JSON(http.StatusOK, resp)

		return
	}

	log.WithFields(logrus.Fields{
		"email": userCreateRequest.Email,
	}).Info("Request Parameters.")

	var localAuth LocalAuth
	err := api.DB.Where("email = ?", userCreateRequest.Email).First(&localAuth).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 用户邮箱没被注册
			localAuth := &LocalAuth{
				Email: userCreateRequest.Email,
			}

			if userCreateRequest.Password == userCreateRequest.PasswordConfirmation {
				// 用户密码两次是一致的情况
				err := localAuth.HashPassword(userCreateRequest.Password)
				if err != nil {
					// 用户密码 hash 过程错误
					log.WithFields(logrus.Fields{
						"error": err,
					}).Error("Hash User Password Error.")

					resp.Success = false
					resp.Code = web.SYS_ERR
					resp.Message = "系统错误，请联系管理员"

					c.JSON(http.StatusOK, resp)
					return
				}

				user := User{
					Name:      localAuth.Email,
					LocalAuth: *localAuth,
				}

				err = api.DB.Create(&user).Error
				if err != nil {
					log.WithFields(logrus.Fields{
						"error": err,
					}).Error("Create User Error.")

					resp.Success = false
					resp.Code = web.SYS_ERR
					resp.Message = "用户创建失败"

					c.JSON(http.StatusOK, resp)
					return
				}

				resp.Success = true
				resp.Code = web.SUCCESS
				resp.Message = "用户创建成功"

				c.JSON(http.StatusOK, resp)

				return

			}

			// 用户密码不一致
			resp.Success = false
			resp.Code = web.SYS_ERR
			resp.Message = "用户两次密码不一致"

			c.JSON(http.StatusOK, resp)

			return
		} else {
			// 查询邮箱出错
			log.WithFields(logrus.Fields{
				"email": userCreateRequest.Email,
				"error": err,
			}).Error("Register User Email Error.")

			resp.Success = false
			resp.Code = web.SYS_ERR
			resp.Message = "系统错误，请联系管理员"

			c.JSON(http.StatusOK, resp)
			return

		}

	} else {
		// 邮箱已存在
		resp.Success = false
		resp.Code = web.SYS_ERR
		resp.Message = "用户邮箱已被注册"

		c.JSON(http.StatusOK, resp)

	}

}
