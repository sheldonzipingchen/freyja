// request.go userAPI请求参数定义
package user

// findUserRequest 用户登录请求
type findUserRequest struct {
	Email    string `json:"email" binding:"required" label:"用户邮箱"`
	Password string `json:"password" binding:"required" label:"密码"`
}
