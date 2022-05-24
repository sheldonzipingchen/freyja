// router.go API 路由配置
package server

import (
	"freyja/db"
	"freyja/lg"
	"freyja/middleware"
	"freyja/rd"
	"freyja/user"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

func newRouter() *gin.Engine {
	log := lg.GetLog()

	r := gin.New()
	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册翻译器
		_ = zh_translations.RegisterDefaultTranslations(v, trans)

		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := fld.Tag.Get("label")
			return name
		})
		log.Info("Setup Validator Engine success.")

	} else {
		log.Warn("Setup Validator Engine error.")

	}

	r.Use(middleware.LoggerMiddleware())
	r.Use(gin.Recovery())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello world")
	})

	apiGroup := r.Group("/api")
	apiV1Group := apiGroup.Group("/v1")

	user.UserAPP(db.GetDB(), rd.GetRDB(), apiV1Group)
	return r
}
