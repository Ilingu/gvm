package gvm

import (
	"fmt"
	"gvm-windows/gvm/utils"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type goDownloader struct {
	version string
}

func MakeGoDownloader(v string) goDownloader {
	return goDownloader{version: v}
}

func (v goDownloader) DownloadMSI() (string, bool) {
	goVersionLink := utils.GenerateWinDownloadUrl(v.version)
	resp, err := http.Get(goVersionLink) // Get from Go Official Website
	if err != nil {
		return "", false
	}
	defer resp.Body.Close()

	// Create temporary file container
	file, err := ioutil.TempFile("", fmt.Sprintf("go%s-*.msi", v.version))
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

func (v goDownloader) DownloadSource() (string, bool) {
	goVersionLink := utils.GenerateSourceDownloadUrl(v.version)
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

	file, err := os.Create(appFolder + fmt.Sprintf("/go%s.src.tar.gz", v.version))
	if os.IsExist(err) {
		log.Println("❌ This Go Version is already installated")
		return "", false
	} else if err != nil {
		log.Println(err)
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
