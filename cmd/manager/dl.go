/*
Copyright © 2022 Ilingu <ilingu@protonmail.com>
*/

package manager

import (
	"fmt"
	"gvm/cmd/cli_helpers"
	"gvm/console"
	"gvm/gvm"

	"github.com/spf13/cobra"
)

// dlCmd represents the dl command
var dlCmd = &cobra.Command{
	Use:   "dl",
	Short: "➡️ Downloads the specified Go MSI Version",
	Long:  "➡️ Downloads the specified Go MSI Version in the app 'AppData' Dir, for later use.",
	Run: func(cmd *cobra.Command, args []string) {
		if !cli_helpers.IsArgsValids(args) {
			cmd.Help()
			return
		}
		version := args[0]

		// Check latest Version
		if version == "latest" {
			latestVersion, ok := cli_helpers.GetLatestGoVersion()
			if !ok {
				console.Error("❌ Cannot fetch latest Go Version. Check your internet connection!")
				return
			}
			version = latestVersion
		}

		// Downloading go
		console.Log(fmt.Sprintf("Downloading go%s... ⏳\n", version))
		Godl := gvm.MakeGoDownloader(version)
		dlPath, err := Godl.DownloadInCache()
		if err != nil {
			console.Error(fmt.Sprintf("Failed to download go%s ❌\n", version))
			return
		}
		console.Success(fmt.Sprintf("go%s Downloaded Successfully ✅: %s\n", version, dlPath))
		console.Log(fmt.Sprintf("You can now execute `gvm manager use %s` to enable this version\n", version))
	},
}

func init() {
}
