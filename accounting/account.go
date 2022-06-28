package accounting

import (
	"gorm.io/gorm"
)

// Account 账户模型
type Account struct {
	gorm.Model
	Name          string // 账户名称
	Code          string // 账号编码
	FullCode      string // 完整账号编码
	Type          string // 账户类型
	IsBankAccount bool   // 是否是银行账户
}
