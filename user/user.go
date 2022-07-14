// user.go 用户基础模型
package user

import (
	"fmt"
	"freyja/utils"

	"gorm.io/gorm"
)

// User 用户基础模型定义
type User struct {
	gorm.Model
	Name      string    `gorm:"comment:'用户昵称'"`
	LocalAuth LocalAuth `gorm:"comment:'用户密码验证信息'"`
}

// TableName 用户基础模型 schema 及表名
func (User) TableName() string {
	return fmt.Sprintf("%s.%s", utils.FREYJA, "users")
}
