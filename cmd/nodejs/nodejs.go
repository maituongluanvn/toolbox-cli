package nodejs

import (
	"fmt"

	"github.com/spf13/cobra"
)

// GitCmd represents the git command
var NodeJS = &cobra.Command{
	Use:   "nodejs",
	Short: "The package contains a NodeJS helper. Use the --help command to view the available commands.",
	Long:  "The package contains a NodeJS helper. Use the --help command to view the available commands.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("The package contains a NodeJS helper. Use the --help command to view the available commands.")
	},
}

func init() {
	// No need to add commands here, it's done in the rootCmd
	// RootCmd.AddCommand(GitCmd)
}
