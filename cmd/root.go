package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "th",
	Short: "Small terraform wrapper",
	Long: `This wrapper check existence of terraform, check his version, run pre-apply hook and post-apply hooks,
	all existing hooks you can found in documentation`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// func init() {
// 	fmt.Println("Init")
// }
