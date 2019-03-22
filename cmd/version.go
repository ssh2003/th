package "version"

import (
    "fmt"
    "github.com/spf13/cobra"
  )

  var versionCmd = &cobra.Command{
  	Use:   "version",
  	Short: "ver",
  	Long:  `This subcommand output current version`,
  	Run: func(cmd *cobra.Command, args []string) {
  		fmt.Println("output version")
  	},
  }

  func init() {
  	RootCmd.AddCommand(versionCmd)
  }
