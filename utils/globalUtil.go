package utils

import (
	"fmt"
	appos "gvm/app_os"
	"net/url"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

// Check if input is a string and then check if input is an empty string
func IsEmptyString(str any) bool {
	realStr, isStr := str.(string)
	return !isStr || len(strings.TrimSpace(realStr)) <= 0
}

// Return whether the input url is a valid HTTP url or not.
func IsValidURL(urlToCheck string) bool {
	_, err := url.ParseRequestURI(urlToCheck)
	return err == nil
}

// Execute a command and returns it output
func ExecCmdWithStdOut(cmd *exec.Cmd) (string, error) {
	res, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// return true if the calling environnement is a test env
func IsTestEnv() bool {
	return os.Getenv("TEST") == "1"
}

// Gets the User's GOROOT. It uses `go env GOROOT` to get that.
func GetGOROOT() (string, bool) {
	goroot, err := ExecCmdWithStdOut(exec.Command("/usr/local/go/bin/go", "env", "GOROOT"))
	if err != nil || goroot == "" {
		return "", false
	}

	return strings.ReplaceAll(goroot, "\n", ""), true
}

// Gets the User's Current Installed version of Go. It uses `go env GOVERSION` to get that. Note that it returns the version in this format: `gox.y.z`, e.g: "go1.19" or "go1.18.5"...
func GetUserGoVersion() (string, bool) {
	goversion, err := ExecCmdWithStdOut(exec.Command("/usr/local/go/bin/go", "env", "GOVERSION"))
	if err != nil || goversion == "" {
		return "", false
	}

	return strings.ReplaceAll(goversion, "\n", ""), true
}

func IsExecutedAsRoot() bool {
	return os.Getuid() == 0
}

// returns the User's Home Directory (%USERPROFILE%)
func GetUserDir() (string, error) {
	sessionUser, err := user.Current()
	if err != nil {
		return "", err
	}

	return sessionUser.HomeDir, nil
}

// returns the CLI directory to cache Go version
func GenerateAppDataPath() (string, error) {
	homeDir, err := GetUserDir()
	if err != nil {
		return "", err
	}

	var appData string
	err = appos.ExecAccording(
		func() {
			appData = fmt.Sprintf("%s/.cache/gvm", homeDir)
		},
		func() {
			appData = fmt.Sprintf("%s/AppData/Roaming/gvm", homeDir)
		},
	)
	if err != nil {
		return "", err
	}

	err = os.MkdirAll(appData, os.ModePerm)
	if err != nil {
		return "", err
	}

	return appData, nil
}
