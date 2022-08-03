/*
Copyright © 2022 Ilingu <ilingu@protonmail.com>
*/

package manager

import (
	"gvm-windows/cmd/utils"
	"gvm-windows/gvm"
	"log"

	"github.com/spf13/cobra"
)

// dlCmd represents the dl command
var dlCmd = &cobra.Command{
	Use:   "dl",
	Short: "➡️ Downloads the specified Go MSI Version",
	Long:  "➡️ Downloads the specified Go MSI Version in the app 'AppData' Dir, for later use.",
	Run: func(cmd *cobra.Command, args []string) {
		if !utils.IsArgsValids(args) {
			cmd.Help()
			return
		}
		version := args[0]

		// Check latest Version
		if version == "latest" {
			latestVersion, ok := utils.GetLatestGoVersion()
			if !ok {
				log.Println("❌ Cannot fetch latest Go Version. Check your internet connection!")
				return
			}
			version = latestVersion
		}

		// Downloading go
		log.Printf("Downloading go%s... ⏳\n", version)
		Godl := gvm.MakeGoDownloader(version)
		dlPath, ok := Godl.DownloadMSI()
		if !ok {
			log.Printf("Failed to download go%s ❌\n", version)
			return
		}
		log.Printf("go%s Downloaded Successfully ✅: %s\n", version, dlPath)
		log.Printf("You can now execute `gvm manager use %s` to enable this version\n", version)
	},
}

func init() {
	// goroot := *dlCmd.Flags().String("goroot", `C:\Program Files\Go`, "Custom GOROOT path, default: C:\\Program Files\\Go ")
}
