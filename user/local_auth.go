// local_auth.go 用户密码验证模型
package user

import (
	"fmt"
	"freyja/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// LocalAuth 用户密码验证信息
type LocalAuth struct {
	gorm.Model
	Email        string `gorm:"comment:'用户邮箱'"`
	PasswordHash string `gorm:"comment:'用户密码'"`
	UserID       uint   `gorm:"comment:'关联到 User 的外键'"`
}

// TableName 用户密码验证信息 schema 及表名
func (LocalAuth) TableName() string {
	return fmt.Sprintf("%s.%s", utils.FREYJA, "local_auth")
}

// HashPassword 对密码明文进行Hash运算
func (auth *LocalAuth) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	auth.PasswordHash = string(bytes)
	return err
}

// CheckPasswordHash 校验密码
func (auth *LocalAuth) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(auth.PasswordHash), []byte(password))
	return err == nil
}
