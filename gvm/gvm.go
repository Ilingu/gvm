package gvm

import (
	appos "gvm/app_os"
	linux_gvm "gvm/gvm/linux"
	windows_gvm "gvm/gvm/windows"
)

func MakeGoDownloader(v string) (out appos.GoDownloader) {
	appos.ExecAccording(
		func() {
			out = linux_gvm.MakeGoDownloader(v)
		},
		func() {
			out = windows_gvm.MakeGoDownloader(v)
		},
	)
	return
}

func MakeGoInstaller(path, version string) (out appos.GoInstaller) {
	appos.ExecAccording(
		func() {
			out = linux_gvm.MakeGoInstaller(path, version)
		},
		func() {
			out = windows_gvm.MakeGoInstaller(path, version)
		},
	)
	return
}
