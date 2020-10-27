package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string
var Verbose bool

var rootCmd = &cobra.Command{
	Use:   "./template",
	Short: "template是常用组件的集合，供参考使用",
	Long:  "template是常用组件的集合，供参考使用",
	Run: func(cmd *cobra.Command, args []string) {
		rootRun()
	},
	Version: "0.0.1",
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func rootRun() {
	fmt.Println("dddddd")
}
