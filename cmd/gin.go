package cmd

import (
	"template/loader"

	"github.com/spf13/cobra"
)

var configPath string

var execCmd = &cobra.Command{
	Use:   "gin",
	Short: "启动gin服务",
	Long:  "使用gin框架，提供http服务，提供的接口有：",
	Run: func(cmd *cobra.Command, args []string) {
		runGin()
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
}

func runGin() {
	loader.LoadConfig(configPath)
}
