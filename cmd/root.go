package cmd

import (
	"bst/pkg/log"
	"github.com/spf13/cobra"
	"os"
)

var configFilePath string

var rootCmd = &cobra.Command{
	SilenceUsage: true,
}

func init() {
	rootCmd.PersistentFlags().StringVar(&configFilePath, "config", "./configs/local.json", "path to json config file")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.New().Error(err)
		os.Exit(1)
	}
}
