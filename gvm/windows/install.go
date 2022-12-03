package windows_gvm

import (
	appos "gvm/app_os"
	"os/exec"
)

type MsiInstaller struct {
	Path    string
	Version string
}

func MakeGoInstaller(p, v string) appos.GoInstaller {
	return MsiInstaller{p, v}
}

// install a Go msi file in the user machine
func (it MsiInstaller) Install() error {
	// `msiexec.exe /i <path_to_msi> /passive` --> will install the msi without popup
	// https://docs.microsoft.com/fr-fr/windows-server/administration/windows-commands/msiexec --> Docs
	return exec.Command("msiexec.exe", "/i", it.Path, "/passive").Run()
}

func (it MsiInstaller) GetPath() string {
	return it.Path
}
