/// response.go 用户API响应定义
package user

/// LoginResponseEntity 用户登录响应
type LoginResponseEntity struct {
	AccessToken string `json:"access_token"`
}
