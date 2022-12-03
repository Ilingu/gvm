package corehelpers

import (
	"fmt"
	appos "gvm/app_os"
	utils "gvm/utils"
	"io"
	"os"
)

/* The tests of these helper are in the app's integration test (yea, I was lazy to do unit tests...) */

// Takes a file buffer and write it into the temp os directory given it's filename (e.g: go-1.19.msi), it returns the created filename
func SaveInTemp(Body io.ReadCloser, filename string) (string, error) {
	// Create temporary file container
	file, err := os.CreateTemp("", filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Populate file with the Go Executable
	_, err = io.Copy(file, Body)
	if err != nil {
		os.Remove(file.Name()) // Remove Temp File
		return "", err
	}

	return file.Name(), nil
}

// Takes a file buffer and write it into the app's cache directory given it's filename (e.g: go-1.19.msi), it returns the created filename
func SaveInCache(Body io.ReadCloser, version string) (string, error) {
	// Create go version in app files
	appFolder, err := utils.GenerateAppDataPath()
	if err != nil {
		return "", err
	}

	var fileDst = appFolder
	err = appos.ExecAccording(
		func() { fileDst += fmt.Sprintf("/go%s.tar.gz", version) },
		func() { fileDst += fmt.Sprintf("/go%s.msi", version) },
	)
	if err != nil {
		return "", err
	}

	file, err := os.Create(fileDst)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Populate file with the Go Executable
	if _, err = io.Copy(file, Body); err != nil {
		os.Remove(fileDst) // Remove File
		return "", err
	}

	return file.Name(), nil
}
