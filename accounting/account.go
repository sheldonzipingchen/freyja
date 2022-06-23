package accounting

import "gorm.io/gorm"

// Account 账户模型
type Account struct {
	gorm.Model
	Name     string
	Code     string
	FullCode string
}
