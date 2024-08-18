package nodejs

import (
	"fmt"

	"github.com/spf13/cobra"
)

// GitCmd represents the git command
var NodeJS = &cobra.Command{
	Use:   "nodejs",
	Short: "Contains git-related helper commands",
	Long:  "A longer description that spans multiple lines and likely contains examples and usage of your command.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Nodejs called")
	},
}

func init() {
	// No need to add commands here, it's done in the rootCmd
	// RootCmd.AddCommand(GitCmd)
}
