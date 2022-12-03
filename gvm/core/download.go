package core

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
)

// Check if URL is valid and return HTTP/GET result
func DownloadGoVersion(URL string) (*http.Response, error) {
	parsedURL, err := url.ParseRequestURI(URL)
	if err != nil {
		return nil, errors.New("invalid url")
	}
	if parsedURL.Host != "go.dev" || parsedURL.Scheme != "https" {
		return nil, errors.New("invalid download url")
	}

	if !strings.Contains(parsedURL.Path, "/dl/") || (!strings.Contains(parsedURL.Path, "linux") && !strings.Contains(parsedURL.Path, "windows")) {
		return nil, errors.New("invalid os")
	}

	return http.Get(URL) // Get from Go Official Website
}
