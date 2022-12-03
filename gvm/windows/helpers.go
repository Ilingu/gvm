package windows_gvm

import (
	"fmt"
	"gvm/gvm/core"
	"net/http"
)

// Generate the download url of the Go MSI
func generateWinDownloadUrl(v string) string {
	return fmt.Sprintf("https://go.dev/dl/go%s.windows-amd64.msi", v)
}

func downloadWinVersion(v string) (*http.Response, error) {
	goVersionLink := generateWinDownloadUrl(v)
	return core.DownloadGoVersion(goVersionLink)
}
