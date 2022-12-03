/*
Copyright © 2022 Ilingu <ilingu@protonmail.com>
*/

package cmd

import (
	"fmt"
	appos "gvm/app_os"
	cobra_helpers "gvm/cmd/cli_helpers"
	"gvm/console"
	"gvm/gvm"
	"gvm/utils"
	"os"

	"github.com/spf13/cobra"
)

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "➡️ Let you switch of Go version easily (recommended).",
	Long:  "➡️ Let you switch of Go Main Version easily: uses the Go MSI executable (recommended for windows).\nIt Downloads the msi executable from Go Official Website if not already downloaded, then it uninstalls the current Go version and finally it installs the newly/already downloaded Go version.\nTIP: use 'latest' arg to download latest go version",
	Run: func(cmd *cobra.Command, args []string) {
		if !cobra_helpers.IsArgsValids(args) {
			cmd.Help()
			return
		}

		targetVersion := args[0]
		UserGoVersion, _ := utils.GetUserGoVersion()

		latestVersion, ok := cobra_helpers.GetLatestGoVersion()
		if targetVersion == "latest" && !ok {
			console.Error("❌ Cannot fetch latest Go Version. Check your internet connection!")
			return
		}

		if targetVersion == "latest" {
			targetVersion = latestVersion
		} else if UserGoVersion != "go"+latestVersion {
			console.Warn(fmt.Sprintf("❕ Go's latest version is %s, you're currently in %s", latestVersion, UserGoVersion))
		}

		// Check Version
		if UserGoVersion == "go"+targetVersion {
			console.Error(fmt.Sprintf("You are already on go%s ❌\n", targetVersion))
			return
		}

		// Get Cache
		appFolder, err := utils.GenerateAppDataPath()
		if err != nil {
			console.Error("You have no Home path")
			return
		}

		var GoCachePath string

		Godl := gvm.MakeGoDownloader(targetVersion)
		var GoInstaller appos.GoInstaller

		err = appos.ExecAccording(
			func() { GoCachePath = appFolder + fmt.Sprintf("/go%s.tar.gz", targetVersion) }, // Linux
			func() { GoCachePath = appFolder + fmt.Sprintf("/go%s.msi", targetVersion) },    // Windows
		)
		if err != nil {
			console.Error(err.Error())
			return
		}

		if temp {
			console.Log(fmt.Sprintf("Downloading go%s... ⏳ (no-cache=%t)\n", targetVersion, temp))
			GoInstaller, err = Godl.DownloadInTemp()
			if err != nil {
				console.Error(fmt.Sprintf("Failed to download go%s ❌\n", targetVersion))
				return
			}
			console.Success(fmt.Sprintf("go%s Downloaded Successfully ✅: %s\n", targetVersion, GoCachePath))
		}

		// Check if file cached, if not download it
		fileInfo, err := os.Stat(GoCachePath)
		if !temp && (os.IsNotExist(err) || err != nil || fileInfo.IsDir()) {
			console.Warn("This Go Version is not downloaded on your machine!")
			console.Log(fmt.Sprintf("Downloading go%s... ⏳ (no-cache=%t)\n", targetVersion, temp))

			GoInstaller, err = Godl.DownloadInCache()
			if err != nil {
				console.Error(fmt.Sprintf("Failed to download go%s ❌, %s\n", targetVersion, err))
				return
			}
			console.Success(fmt.Sprintf("go%s Downloaded Successfully ✅: %s\n", targetVersion, GoCachePath))
		} else if !temp {
			GoInstaller = gvm.MakeGoInstaller(GoCachePath, targetVersion)
		}

		console.Log(fmt.Sprintf("Installing go%s... ⏳\n", targetVersion))
		if GoInstaller == nil {
			console.Error("No default installer, failed")
		}

		err = GoInstaller.Install()
		if err != nil {
			console.Error(fmt.Sprintf("Failed to install go%s ❌\n", targetVersion))
			os.Remove(GoCachePath) // Remove Corrupted File
			return
		}

		console.Success(fmt.Sprintf("go%s Installed Successfully ✅\n", targetVersion))
	},
}

var temp bool

func init() {
	rootCmd.AddCommand(switchCmd)
	switchCmd.Flags().BoolVar(&temp, "no-cache", false, "Whether or not the downloaded Go Version will be cached in disk")
}
