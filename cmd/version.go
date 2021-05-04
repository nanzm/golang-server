package cmd

import (
	"dora/config"
	"dora/pkg/logger"
	"github.com/spf13/cobra"
)

var (
	build   = config.Build
	compile = config.Compile
	version = config.Version
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Println("Build:" + build)
		logger.Println("Compile:" + compile)
		logger.Println("Version:" + version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
