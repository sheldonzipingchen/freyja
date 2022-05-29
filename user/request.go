// request.go userAPI请求参数定义
package user

// findUserRequest 用户登录请求
type findUserRequest struct {
	Email    string `json:"email" binding:"required" label:"用户邮箱"`
	Password string `json:"password" binding:"required" label:"密码"`
}

// userCreateRequest 用户创建请求
type userCreateRequest struct {
	Email                string `json:"email" binding:"required,email" label:"用户邮箱"`
	Password             string `json:"password" binding:"required" label:"密码"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required" label:"确认密码"`
}
