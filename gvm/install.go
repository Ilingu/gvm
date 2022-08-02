package gvm

import (
	"gvm-windows/gvm/utils"
	"log"
	"os"
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
	// https://docs.microsoft.com/fr-fr/windows-server/administration/windows-commands/msiexec --> Docs
	err := exec.Command("msiexec.exe", "/i", it.path, "/passive").Run()
	return err == nil
}

func (it goInstaller) InstallAsSource() bool {
	err := os.Chdir(it.path + "\\src") // Change dir to be inside Go Installer
	if err != nil {
		return false
	}

	output, err := utils.ExecCmdWithStdOut(exec.Command("all.bat"))
	if err != nil {
		log.Println(err)
		return false
	}
	log.Println(output)

	return true
}
