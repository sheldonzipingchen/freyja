package commands

import (
	"freyja/server"

	"github.com/spf13/cobra"
)

var apiServerCommand = &cobra.Command{
	Use:   "apiserver",
	Short: "启动 API Server",
	Long:  `启动 API Server`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Init()
	},
}
