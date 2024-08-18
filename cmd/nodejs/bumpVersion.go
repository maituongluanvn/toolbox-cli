package nodejs

import (
	"fmt"
	"log"
	packagejson "toolbox-cli/package"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// bumpVersionCmd represents the bumpVersion command
var bumpVersionCmd = &cobra.Command{
	Use:   "bump-version",
	Short: "Use for bump version on package.json file",
	Long: `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get version from package.json
		version, err := packagejson.GetVersion("package.json")
		if err != nil {
			log.Fatalf("Failed to get version, please check your package.json: %v", err)
		}
		fmt.Printf("Now package version: %s\n", version)

		// Define the question
		prompt := &survey.Input{
			Message: "Which version do you want to bump up to ?",
		}

		// Store the answer
		var newVersion string

		// Ask the question about version
		errAsk := survey.AskOne(prompt, &newVersion)
		if errAsk != nil {
			fmt.Println("Failed to get input:", err)
			return
		}


		packagejson.BumpVersion("package.json",newVersion)

		// Print the answer
		fmt.Printf("Hello, %s!\n", newVersion)
	},
}

func init() {
	NodeJS.AddCommand(bumpVersionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bumpVersionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bumpVersionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
