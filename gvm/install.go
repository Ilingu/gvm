package gvm

import (
	"fmt"
	"gvm-windows/gvm/utils"
	"log"
	"os"
	"os/exec"
)

type goInstaller struct {
	path    string
	version string
}

func MakeGoInstaller(path, v string) goInstaller {
	return goInstaller{path: path, version: v}
}

func (it goInstaller) InstallAsMSI() bool {
	// `msiexec.exe /i <path_to_msi> /passive` --> will install the msi without popup
	// https://docs.microsoft.com/fr-fr/windows-server/administration/windows-commands/msiexec --> Docs
	err := exec.Command("msiexec.exe", "/i", it.path, "/passive").Run()
	return err == nil
}

func (it goInstaller) InstallAsSource() bool {
	// defer os.Remove(it.path) // Remove first
	appDir, err := utils.GenerateAppDataPath()
	if err != nil {
		return false
	}

	GoRootFolder := appDir + fmt.Sprintf("/go%s", it.version)
	err = os.Mkdir(GoRootFolder, os.ModePerm)
	if !os.IsExist(err) && err != nil {
		return false
	}

	err = utils.Untar(it.path, GoRootFolder) // Untar the file, and put it in the right dir
	if err != nil {
		log.Println(err)
		return false
	}

	// utils.ExecCmdWithStdOut(exec.Command(""))

	return true
}
