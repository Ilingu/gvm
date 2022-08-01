package utils

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
)

func GetUserDir() (string, error) {
	sessionUser, err := user.Current()
	if err != nil {
		return "", err
	}

	return sessionUser.HomeDir, nil
}

func GenerateAppDataPath() (string, error) {
	homeDir, err := GetUserDir()
	if err != nil {
		return "", err
	}

	appData := fmt.Sprintf("%s/AppData/Roaming/gvm-windows", homeDir)
	err = os.MkdirAll(appData, os.ModePerm)
	if err != nil {
		return "", err
	}

	return appData, nil
}

func GenerateDownloadUrl(v string) string {
	return fmt.Sprintf("https://go.dev/dl/go%s.windows-amd64.msi", v)
}

func ExecCmd(cmd *exec.Cmd) error {
	err := cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	return err
}
