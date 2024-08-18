package packagejson

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/blang/semver/v4"
	"github.com/tidwall/sjson"
)

// PackageJSON struct to map the package.json structure
type PackageJSON struct {
	Version string `json:"version"`
}

// GetVersion reads the version from the specified package.json file
func GetVersion(filePath string) (string, error) {
	// Open package.json file
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read the file content
	byteValue, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Unmarshal JSON into PackageJSON struct
	var packageJSON PackageJSON
	err = json.Unmarshal(byteValue, &packageJSON)
	if err != nil {
		return "", err
	}

	// Return the version
	return packageJSON.Version, nil
}

// BumpVersion updates the version in the specified package.json file
func BumpVersion(filePath, newVersion string) error {
	// Validate the new version
	_, errVersion := semver.Parse(newVersion)
	if errVersion != nil {
		return errVersion
	}

	// Read the existing file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Update the version field using sjson to keep the position
	updatedContent, err := sjson.SetBytes(content, "version", newVersion)
	if err != nil {
		return err
	}

	// Convert the updated content back to a string and check if the version is still at line 3
	lines := strings.Split(string(updatedContent), "\n")
	if len(lines) < 3 || !strings.Contains(lines[2], `"version"`) {
		return errors.New("failed to maintain version position")
	}

	// Write the updated content back to the file
	err = os.WriteFile(filePath, updatedContent, 0644)
	if err != nil {
		return err
	}

	return nil
}
