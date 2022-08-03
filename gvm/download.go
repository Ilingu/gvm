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

func MakeGoDownloader(v string) goDownloader {
	return goDownloader{version: v}
}

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

	file, err := os.Create(appFolder + fmt.Sprintf("\\go%s.msi", v.version))
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
	// Extract File
	// defer os.Remove(path) // Second Remove .tar.gz file
	// defer file.Close()    // First Close .tar.gz file

	// GoRootFolder := appFolder + fmt.Sprintf("\\go%s", v.version)
	// err = os.Mkdir(GoRootFolder, os.ModePerm)
	// if !os.IsExist(err) && err != nil {
	// 	return "", false
	// }

	// log.Println("Almost There! Extracting Go Files...")
	// err = utils.Untar(path, GoRootFolder) // Untar the file, and put it in the right dir
	// if err != nil {
	// 	return "", false
	// }
}
