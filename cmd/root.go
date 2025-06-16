package cmd

import (
	"github.com/spf13/cobra"
)

var (
	confFilePath = ""
	rootCmd      = &cobra.Command{
		Use:          "start",
		SilenceUsage: true,
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&confFilePath, "config", "c", "server.yaml", "config file path")
	rootCmd.AddCommand(startCmd)
}

func Execute() {
	_ = rootCmd.Execute()
}
