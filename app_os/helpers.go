package appos

import (
	"errors"
	"runtime"
)

type OperatingSystem int64

const (
	LINUX   OperatingSystem = iota
	WINDOWS OperatingSystem = iota
)

type GoDownloader interface {
	DownloadInTemp() (GoInstaller, error)
	DownloadInCache() (GoInstaller, error)
}

type GoInstaller interface {
	Install() error
	GetPath() string
}

// Return the current user OS
func Which() OperatingSystem {
	switch runtime.GOOS {
	case "linux":
		return LINUX
	case "windows":
		return WINDOWS
	default:
		return -1
	}

}

// Execute the piece of code corresponding to the user os
func ExecAccording(linux, windows func()) error {
	switch Which() {
	case LINUX:
		linux()
		return nil
	case WINDOWS:
		windows()
		return nil
	}
	return errors.New("os not supported")
}

// Return if the cli app can handle this os
func IsSupported() bool {
	return Which() != -1
}
