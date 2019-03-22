package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "io/ioutil"

  )
//var flagvar int
var runCmd = &cobra.Command{
  	Use:   "run",
  	Short: "run terraform apply with all pre and post actions",
  	Long:  `This subcommand run pre-hook, terraform init, terraform apply and post-hook`,
  	Run: func(cmd *cobra.Command, args []string) {
  		fmt.Printf("Run\n")
      fmt.Println("config file name ", configFileName)
      dat, err := ioutil.ReadFile(configFileName)
      if err != nil {
        panic(err)
      }
      fmt.Print(string(dat))
  	},
  }

  func init() {
  	RootCmd.AddCommand(runCmd)
  }
