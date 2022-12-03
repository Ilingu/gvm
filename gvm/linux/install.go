package linux_gvm

import (
	"errors"
	"fmt"
	appos "gvm/app_os"
	console "gvm/console"
	utils "gvm/utils"
	"os/exec"
)

type LinuxInstaller struct {
	Path    string
	Version string
}

func MakeGoInstaller(p, v string) appos.GoInstaller {
	return LinuxInstaller{p, v}
}

func (it LinuxInstaller) GetPath() string {
	return it.Path
}

// install a Go linux file in the user machine
func (it LinuxInstaller) Install() error {
	_, exist := utils.GetUserGoVersion()
	if exist {
		console.Log("⏳ Removing old version of Go...")

		if !utils.IsExecutedAsRoot() {
			return errors.New("no root access, please retry with `sudo`")
		}
		if err := exec.Command("rm", "-rf", "/usr/local/go").Run(); err != nil {
			return errors.New("no root access, please retry with `sudo`")
		}
	}

	console.Log("⏳ Unpack new version and install it...")
	if err := untar(it.Path, "/usr/local/go"); err != nil {
		return errors.New("couldn't untar download file")
	}

	if !exist {
		console.Log("🔎 Setting Path...")
		if err := exec.Command("export", "PATH=$PATH:/usr/local/go/bin").Run(); err != nil {
			return errors.New("couldn't export path")
		}
	}

	console.Log("⏳ Checking installation...")
	if newGoVersion, ok := utils.GetUserGoVersion(); !ok || newGoVersion != "go"+it.Version {
		return fmt.Errorf("failed to install go%s", newGoVersion)
	}

	return nil
}
