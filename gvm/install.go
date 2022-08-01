package gvm

import (
	"gvm-windows/gvm/utils"
	"os/exec"
)

type goInstaller struct {
	path string
}

func MakeGoInstaller(path string) goInstaller {
	return goInstaller{path: path}
}

func (it goInstaller) InstallAsMSI() bool {
	// `msiexec.exe /i <path_to_msi> /passive` --> will install the msi without popup
	err := utils.ExecCmd(exec.Command("msiexec.exe", "/i", it.path, "/passive"))
	return err == nil
}

func (it goInstaller) InstallAsSource() bool {
	return true
}
