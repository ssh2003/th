package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
  )

const Version = "0.0.1"
var versionCmd = &cobra.Command{
  	Use:   "version",
  	Short: "Output verion of th",
  	Long:  `This subcommand output current version`,
  	Run: func(cmd *cobra.Command, args []string) {
  		fmt.Printf("th version %s\n", Version)
  	},
  }

  func init() {
  	RootCmd.AddCommand(versionCmd)
  }
