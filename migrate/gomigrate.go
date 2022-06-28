package migrate

import (
	"freyja/db"

	"github.com/go-gormigrate/gormigrate/v2"
)

var gomigrate *gormigrate.Gormigrate

// Init 初始化对象
func Init() {
	db := db.GetDB()
	gomigrate = gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{})
}

// GetGomigrate 获取gomigrate
func GetGomigrate() *gormigrate.Gormigrate {
	return gomigrate
}
