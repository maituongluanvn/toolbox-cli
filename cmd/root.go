/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "toolbox",
	Short: "A brief description of your application",
	Long: `A longer description`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.toolbox-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	// Config
	// cobra.OnInitialize(initConfig)
  	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("license", "apache")
}


// func initConfig() {
// 	// Don't forget to read config either from cfgFile or from home directory!
// 	if cfgFile != "" {
// 	  // Use config file from the flag.
// 	  viper.SetConfigFile(cfgFile)
// 	} else {
// 	  // Find home directory.
// 	  home, err := homedir.Dir()
// 	  if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	  }
  
// 	  // Search config in home directory with name ".cobra" (without extension).
// 	  viper.AddConfigPath(home)
// 	  viper.SetConfigName(".cobra")
// 	}
  
// 	if err := viper.ReadInConfig(); err != nil {
// 	  fmt.Println("Can't read config:", err)
// 	  os.Exit(1)
// 	}
//   }
