/*
Copyright © 2022 Ilingu <ilingu@protonmail.com>
*/

package manager

import (
	"fmt"
	appos "gvm/app_os"
	"gvm/cmd/cli_helpers"
	"gvm/console"
	"gvm/gvm"
	"gvm/utils"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "➡️ Switch between multiples version of Go",
	Long:  `➡️ Switch between multiples version of Go. If the specified Go Version is not downloaded the process exit.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !cli_helpers.IsArgsValids(args) {
			cmd.Help()
			return
		}

		targetVersion := args[0]
		UserGoVersion, _ := utils.GetUserGoVersion()

		if targetVersion == "latest" {
			latestVersion, ok := cli_helpers.GetLatestGoVersion()
			if !ok {
				console.Error("❌ Cannot fetch latest Go Version. Check your internet connection!")
				return
			}
			targetVersion = latestVersion
		}

		// Check Version
		if strings.Contains(UserGoVersion, targetVersion) {
			console.Error(fmt.Sprintf("You are already on go%s ❌\n", targetVersion))
			return
		}

		appFolder, err := utils.GenerateAppDataPath()
		if err != nil {
			return
		}

		var GoCachePath string

		err = appos.ExecAccording(
			func() { GoCachePath = appFolder + fmt.Sprintf("/go%s.tar.gz", targetVersion) }, // Linux
			func() { GoCachePath = appFolder + fmt.Sprintf("/go%s.msi", targetVersion) },    // Windows
		)
		if err != nil {
			console.Error(err.Error())
			return
		}

		fileInfo, err := os.Stat(GoCachePath)
		if os.IsNotExist(err) || err != nil || fileInfo.IsDir() {
			console.Error(fmt.Sprintf("❌ This Go Version is not downloaded on your machine!\nType: `gvm manager dl %s` to download this version", targetVersion))
			return
		}

		console.Log(fmt.Sprintf("Switching to go%s... ⏳\n", targetVersion))
		GoInstaller := gvm.MakeGoInstaller(GoCachePath, targetVersion)
		if err := GoInstaller.Install(); err != nil {
			console.Error(fmt.Sprintf("Failed to switch to go%s ❌\n", targetVersion))
			return
		}

		console.Success(fmt.Sprintf("Switched to go%s Successfully ✅\n", targetVersion))
	},
}

func init() {
}
