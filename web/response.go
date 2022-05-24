// response.go api 响应相关定义
package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Response 响应模型
type Response struct {
	Success bool        `json:"success"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *Response) String() string {
	return fmt.Sprintf("Success: %v, Code: %v, Message: %v, Data: %v", r.Success, r.Code, r.Message, r.Data)
}

// GetGinResponse 生成Gin响应请求
func (r *Response) GetGinResponse() map[string]interface{} {
	return gin.H{
		"success": r.Success,
		"code":    r.Code,
		"message": r.Message,
		"data":    r.Data,
	}
}
