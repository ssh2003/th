package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	//"github.com/spf13/viper"
)

//var cfgFile string
var configFileName string
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

func init() {
	fmt.Println("Init")
	flag.StringVar(&configFileName, "config", ".tfconfig.json", "config file of wrapper")
  flag.Parse()

	// cobra.OnInitialize(initConfig)
	// RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./.tfconfig.json)")
}

// func initConfig() {
// 	fmt.Println("Init Config")
// 	if cfgFile != "" {
// 		// Use config file from the flag.
// 		viper.SetConfigFile(cfgFile)
// 	} else {
//
// 		// Search config in home directory with name ".cobra-example" (without extension).
// 		viper.AddConfigPath(".")
// 		viper.SetConfigName(".tfconfig.json")
// 	}
//
// 	viper.AutomaticEnv() // read in environment variables that match
//
// 	// If a config file is found, read it in.
// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Println("Using config file:", viper.ConfigFileUsed())
// 	}else{
// 		fmt.Println("Can't read config file: .tfconfig.json", err)
// 	}
// }
