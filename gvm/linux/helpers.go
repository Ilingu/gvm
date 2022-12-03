package linux_gvm

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"gvm/gvm/core"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Generate the download url of Go Source
func generateLinuxDownloadUrl(v string) string {
	return fmt.Sprintf("https://go.dev/dl/go%s.linux-amd64.tar.gz", v)
}

func downloadLinuxVersion(v string) (*http.Response, error) {
	goVersionLink := generateLinuxDownloadUrl(v)
	return core.DownloadGoVersion(goVersionLink)
}

// untar takes a destination path and a reader; a tar reader loops over the tarfile
// creating the file structure at 'dst' along the way, and writing any files
func untar(from, dst string) error {
	reader, err := os.Open(from)
	if err != nil {
		return err
	}
	defer reader.Close()

	// UnGzip
	gzr, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)
	for {
		header, err := tr.Next()

		switch {
		case err == io.EOF: // if no more files are found return
			return nil

		case err != nil: // return any other error
			return err

		case header == nil: // if the header is nil, just skip it (not sure how this happens)
			continue
		}

		// the target location where the dir/file should be created
		target := filepath.Join(dst, strings.Replace(header.Name, "go/", "", 1))

		// check the file type
		switch header.Typeflag {
		// if its a dir and it doesn't exist create it
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}

		// if it's a file create it
		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode)) // if there is a root parent, this will not create it and it will result in a error (like: "path doesn't exist") (1/2)
			if err != nil {
				return err // to avoid that you'll have to MkdirAll before (2/2)
			}

			// copy over contents
			if _, err := io.Copy(f, tr); err != nil {
				return err
			}

			// manually close here after each file operation; defering would cause each file close
			// to wait until all operations have completed.
			f.Close()
		}
	}
}
