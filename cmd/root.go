package cmd

import (
	"dora/pkg/logger"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"
)

func Execute() {
	lc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	time.Local = lc

	fmt.Println()
	fmt.Println("                              _     _                                      \n                             ( \\---/ )                                     \n                              ) . . (                                      \n________________________,--._(___Y___)_,--._______________________      \n                        `--'           `--'                                ")
	fmt.Println()

	if err := rootCmd.Execute(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "dora",
	Short: "dora is a platform to help web developer improve efficiency.",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Println("hi. my name is dora!")
	},
}
