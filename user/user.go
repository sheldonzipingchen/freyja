// user.go 用户基础模型
package user

import "gorm.io/gorm"

// User 用户基础模型定义
type User struct {
	gorm.Model
	Name      string    `gorm:"comment:'用户昵称'"`
	LocalAuth LocalAuth `gorm:"comment:'用户密码验证信息'"`
}

func (User) TableName() string {
	return "freyja.users"
}
