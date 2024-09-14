package nodejs

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	packagejson "toolbox-cli/package"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var bumpVersionCmd = &cobra.Command{
	Use:   "bump-version",
	Short: "Use for bump version on package.json file",
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
		if err := survey.AskOne(prompt, &newVersion); err != nil {
			fmt.Println("Failed to get input:", err)
			return
		}

		if errBumpVersion := packagejson.BumpVersion("package.json", newVersion); errBumpVersion != nil {
			fmt.Println("Not correct version format: ", errBumpVersion)
			return
		}

		isCreatePullRequestPrompt := &survey.Input{
			Message: "Do you want to comment then create pull request to Github? (yes/no)",
		}

		// Store the answer
		var createPullRequest string

		// Ask if user wants to create a pull request
		if err := survey.AskOne(isCreatePullRequestPrompt, &createPullRequest); err != nil {
			fmt.Println("Failed to get input:", err)
			return
		}

		// Trim spaces and convert to lowercase for consistent comparison
		createPullRequest = strings.TrimSpace(strings.ToLower(createPullRequest))

		if createPullRequest != "yes" {
			fmt.Println("No pull request will be created.")
			return
		}

		// Create a new GitHub CLI command to create pull request
		prTitle := "Bump version to " + newVersion
		prBody := "This pull request bumps the package version to " + newVersion
		headBranch := "bump-version-to-" + newVersion // Avoid spaces in branch names
		baseBranch := "master" // Changed base branch to master

		// Ensure the branch exists; if not, create it
		if err := createBranchIfNotExists(headBranch); err != nil {
			fmt.Printf("Failed to create branch: %v\n", err)
			return
		}

		// Commit changes with a message
		if err := commitChanges("bump version to " + newVersion); err != nil {
			fmt.Printf("Failed to commit changes: %v\n", err)
			return
		}

		// Push the branch to remote
		if err := pushBranch(headBranch); err != nil {
			fmt.Printf("Failed to push branch to remote: %v\n", err)
			return
		}

		// Create the pull request using `gh` CLI
		ghCreatePrCmd := exec.Command("gh", "pr", "create", "--title", prTitle, "--body", prBody, "--head", headBranch, "--base", baseBranch)
		output, err := ghCreatePrCmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Failed to create pull request: %s\n", err)
			fmt.Println(string(output))
			return
		}

		fmt.Println("Pull request created successfully!")
	},
}

// createBranchIfNotExists creates the specified branch if it does not already exist
func createBranchIfNotExists(branchName string) error {
	// Check if branch already exists
	cmd := exec.Command("git", "show-ref", "--verify", "--quiet", "refs/heads/"+branchName)
	err := cmd.Run()

	if err != nil {
		if strings.Contains(err.Error(), "fatal: Not a git repository") {
			return fmt.Errorf("not a git repository")
		}

		if err.Error() == "exit status 1" {
			// Branch does not exist, create it
			cmd = exec.Command("git", "checkout", "-b", branchName)
			if err := cmd.Run(); err != nil {
				return fmt.Errorf("failed to create branch: %v", err)
			}
		} else {
			return fmt.Errorf("error checking branch existence: %v", err)
		}
	}

	return nil
}

// commitChanges commits the changes with the given message
func commitChanges(message string) error {
	cmd := exec.Command("git", "add", ".")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to add changes: %v", err)
	}

	cmd = exec.Command("git", "commit", "-m", message)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to commit changes: %v", err)
	}

	return nil
}

// pushBranch pushes the branch to remote
func pushBranch(branchName string) error {
	cmd := exec.Command("git", "push", "-u", "origin", branchName)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to push branch to remote: %v", err)
	}

	return nil
}

func init() {
	NodeJS.AddCommand(bumpVersionCmd)
}
