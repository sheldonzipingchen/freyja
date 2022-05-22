package commands

import (
	"freyja/config"
	"freyja/db"
	"freyja/lg"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var environment string

var rootCmd = &cobra.Command{
	Use:   "freyja",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&environment, "environment", "e", "development", "应用运行的模式(development, test, production)")

	// 添加 apiserver 命令
	rootCmd.AddCommand(apiServerCommand)
}

func initConfig() {
	config.Init(environment)
	lg.Init()
	db.Init()
}

func Execute() {
	log := lg.GetLog()
	if err := rootCmd.Execute(); err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Command Execute Fatal.")
	}
}
