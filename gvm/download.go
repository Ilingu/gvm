package gvm

import (
	"fmt"
	"gvm-windows/gvm/utils"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type goDownloader struct {
	version string
}

func MakeGoDownloader(v string) goDownloader {
	return goDownloader{version: v}
}

func (v goDownloader) DownloadMSI() (string, bool) {
	goVersionLink := utils.GenerateDownloadUrl(v.version)
	resp, err := http.Get(goVersionLink) // Get from Go Official Website
	if err != nil {
		return "", false
	}
	defer resp.Body.Close()

	// Create temporary file container
	file, err := ioutil.TempFile("", fmt.Sprintf("go%s-*.msi", v.version))
	if err != nil {
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

func (v goDownloader) DownloadSource() bool {
	return true
}
