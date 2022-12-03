package windows_gvm

import (
	"fmt"
	appos "gvm/app_os"
	corehelpers "gvm/gvm/core/helpers"
)

type MsiDl struct {
	version string
}

// Create a Go Downloader given the wanted version to download, it return a Go<Version>Downloader
func MakeGoDownloader(v string) appos.GoDownloader {
	return MsiDl{version: v}
}

// Download the Go version MSI file from the Go Official Website, but doesn't cache it (stored in "/temp" dir)
func (v MsiDl) DownloadInTemp() (appos.GoInstaller, error) {
	resp, err := downloadWinVersion(v.version) // Get from Go Official Website
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	path, err := corehelpers.SaveInTemp(resp.Body, fmt.Sprintf("go%s-*.msi", v.version))
	if err != nil {
		return nil, err
	}
	return MsiInstaller{path, v.version}, nil
}

// Download the Go version MSI file from the Go Official Website, and cache it in the app directory (for future use)
func (v MsiDl) DownloadInCache() (appos.GoInstaller, error) {
	resp, err := downloadWinVersion(v.version) // Get from Go Official Website
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	path, err := corehelpers.SaveInCache(resp.Body, v.version)
	if err != nil {
		return nil, err
	}
	return &MsiInstaller{path, v.version}, nil
}
