package gvm

import (
	"fmt"
	"gvm-windows/gvm/utils"
	"io"
	"net/http"
	"os"
)

type goDownloader struct {
	version string
}

// Create a Go Downloader given the wanted version to download, it return a Go<Version>Downloader
func MakeGoDownloader(v string) goDownloader {
	return goDownloader{version: v}
}

// Download the Go version MSI file from the Go Official Website, but doesn't cache it (stored in "/temp" dir)
func (v goDownloader) DownloadTempMSI() (string, bool) {
	goVersionLink := utils.GenerateWinDownloadUrl(v.version)
	resp, err := http.Get(goVersionLink) // Get from Go Official Website
	if err != nil {
		return "", false
	}
	defer resp.Body.Close()

	// Create temporary file container
	file, err := os.CreateTemp("", fmt.Sprintf("go%s-*.msi", v.version))
	if err != nil {
		return "", false
	}
	defer file.Close()

	// Populate file with the Go Executable
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", false
	}

	return file.Name(), true
}

// Download the Go version MSI file from the Go Official Website, and cache it in the app directory (for future use)
func (v goDownloader) DownloadMSI() (string, bool) {
	goVersionLink := utils.GenerateWinDownloadUrl(v.version)
	resp, err := http.Get(goVersionLink) // Get from Go Official Website
	if err != nil {
		return "", false
	}
	defer resp.Body.Close()

	// Create go version in app files
	appFolder, err := utils.GenerateAppDataPath()
	if err != nil {
		return "", false
	}

	var fileDst = appFolder
	if utils.IsTestEnv() {
		fileDst += fmt.Sprintf("\\go%s-test.msi", v.version)
	} else {
		fileDst += fmt.Sprintf("\\go%s.msi", v.version)
	}

	file, err := os.Create(fileDst)
	if err != nil {
		return "", false
	}
	defer file.Close()

	// Populate file with the Go Executable
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", false
	}

	return file.Name(), true
}
