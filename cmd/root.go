package cmd

import (
	"fmt"
	"os"
	"toolbox-cli/cmd/nodejs"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Exported variable to be used in other packages
var RootCmd = &cobra.Command{
	Use:   "toolbox-cli",
	Short: "A brief description of your application",
	Long:  `A longer description of your application.`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func addSubcommandsPalettes(){
	RootCmd.AddCommand(nodejs.NodeJS)
}

func init() {
	RootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// author command
	viper.BindPFlag("author", RootCmd.PersistentFlags().Lookup("author"))
	viper.SetDefault("author", "Luan Mai <maituongluanvn@gmail.com>")
	viper.SetDefault("license", "apache")

	addSubcommandsPalettes()
	// Initialize commands
	// initCommands()
}

// func initCommands() {
// 	RootCmd.AddCommand(git.GitCmd) // Add command from git package
// }
