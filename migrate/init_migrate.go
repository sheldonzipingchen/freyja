package migrate

import (
	"freyja/lg"
	"freyja/user"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func InitMigrate() {
	m := GetGomigrate()

	log := lg.GetLog()

	log.Info("Begin Migrate Database.")

	m.InitSchema(func(tx *gorm.DB) error {
		err := tx.AutoMigrate(
			// 初始化 user 包模型
			&user.LocalAuth{},
			&user.User{},
		)

		if err != nil {
			log.WithFields(logrus.Fields{
				"error": err,
			}).Error("InitMigrate Database Error.")
		}
		return err
	})

	log.Info("End Migrate Database.")
}
