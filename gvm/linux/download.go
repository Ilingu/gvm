package linux_gvm

import (
	"fmt"
	appos "gvm/app_os"
	corehelpers "gvm/gvm/core/helpers"
)

type Linuxdl struct {
	version string
}

// Create a Go Downloader given the wanted version to download, it return a Go<Version>Downloader
func MakeGoDownloader(v string) appos.GoDownloader {
	return Linuxdl{version: v}
}

func (lgd Linuxdl) DownloadInTemp() (appos.GoInstaller, error) {
	resp, err := downloadLinuxVersion(lgd.version) // Get from Go Official Website
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	path, err := corehelpers.SaveInTemp(resp.Body, fmt.Sprintf("go%s-*.tar.gz", lgd.version))
	if err != nil {
		return nil, err
	}
	return &LinuxInstaller{path, lgd.version}, nil
}

func (lgd Linuxdl) DownloadInCache() (appos.GoInstaller, error) {
	resp, err := downloadLinuxVersion(lgd.version) // Get from Go Official Website
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	path, err := corehelpers.SaveInCache(resp.Body, lgd.version)
	if err != nil {
		return nil, err
	}
	return &LinuxInstaller{path, lgd.version}, nil
}
