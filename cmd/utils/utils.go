package utils

import (
	"gvm-windows/gvm/utils"
	"io/ioutil"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
)

func IsArgsValids(args []string) bool {
	if len(args) != 1 {
		return false
	}
	if args[0] == "latest" {
		return true
	}

	checkArgShape := regexp.MustCompile(`^[0-9]+\.[0-9]+(?:\.[0-9]+)?$`)
	return checkArgShape.MatchString(args[0])
}

func GetGOROOT() (string, bool) {
	goroot, err := utils.ExecCmdWithStdOut(exec.Command("go", "env", "GOROOT"))
	if err != nil || goroot == "" {
		return "", false
	}

	goroot = strings.Map(func(r rune) rune {
		switch r {
		case '\n':
			return rune(0)
		default:
			return r
		}
	}, goroot)

	return goroot, true
}

func GetLatestGoVersion() (string, bool) {
	resp, err := http.Get("https://go.dev/VERSION?m=text") // Get Version Info from Go Official Website
	if err != nil {
		return "", false
	}
	defer resp.Body.Close()

	RawVersion, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", false
	}

	version := string(RawVersion)
	return strings.Trim(version, "go"), true
}
