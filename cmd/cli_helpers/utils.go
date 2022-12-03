package cli_helpers

import (
	"io"
	"net/http"
	"regexp"
	"strings"
)

// Checks if the args passed to the CLI are valids. Criteria: len(args) == 1 && (args[0] is a valid GoVersion || args[0] == "latest")
func IsArgsValids(args []string) bool {
	if len(args) != 1 {
		return false
	}
	if args[0] == "latest" {
		return true
	}

	checkArgShape := regexp.MustCompile(`^[0-9]+\.[0-9]+(?:\.[0-9]+)?$`)
	return checkArgShape.MatchString(args[0])
}

// It will go get the latest official release of Go from the Go Website (https://go.dev).
// Note that this returns the version in this format: `x.y.z`, e.g: "1.19" or "1.18.5"...
func GetLatestGoVersion() (string, bool) {
	resp, err := http.Get("https://go.dev/VERSION?m=text") // Get Version Info from Go Official Website
	if err != nil {
		return "", false
	}
	defer resp.Body.Close()

	RawVersion, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", false
	}

	version := string(RawVersion)
	return strings.Trim(version, "go"), true
}
