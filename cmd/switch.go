/*
Copyright © 2022 Ilingu <ilingu@protonmail.com>
*/

package cmd

import (
	"errors"
	"fmt"
	"gvm-windows/cmd/utils"
	"gvm-windows/gvm"
	gvmUtils "gvm-windows/gvm/utils"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "➡️ Let you switch of Go version easily (recommended).",
	Long:  "➡️ Let you switch of Go Main Version easily: uses the Go MSI executable (recommended for windows).\nIt Downloads the msi executable from Go Official Website if not already downloaded, then it uninstalls the current Go version and finally it installs the newly/already downloaded Go version.\nTIP: use 'latest' arg to download latest go version",
	Run: func(cmd *cobra.Command, args []string) {
		if !utils.IsArgsValids(args) {
			cmd.Help()
			return
		}
		version := args[0]

		latestVersion, ok := utils.GetLatestGoVersion()
		if version == "latest" && !ok {
			log.Println("❌ Cannot fetch latest Go Version. Check your internet connection!")
			return
		}

		if version == "latest" {
			version = latestVersion
		} else if runtime.Version() != "go"+latestVersion {
			log.Printf("❕ Go's latest version is %s, you're currently in %s", latestVersion, runtime.Version())
		}

		// Check Version
		if strings.Contains(runtime.Version(), version) {
			log.Printf("You are already on go%s ❌\n", version)
			return
		}

		// Get Cache
		appFolder, err := gvmUtils.GenerateAppDataPath()
		if err != nil {
			return
		}

		GoMsiExecutable, ok := appFolder+fmt.Sprintf("\\go%s.msi", version), false
		Godl := gvm.MakeGoDownloader(version)
		downloadGo := func() error {
			log.Printf("Downloading go%s... ⏳ (no-cache=%t)\n", version, temp)

			if temp {
				GoMsiExecutable, ok = Godl.DownloadTempMSI()
			} else {
				GoMsiExecutable, ok = Godl.DownloadMSI()
			}

			if !ok {
				log.Printf("Failed to download go%s ❌\n", version)
				return errors.New("")
			}

			log.Printf("go%s Downloaded Successfully ✅: %s\n", version, GoMsiExecutable)
			return nil
		}

		fileInfo, err := os.Stat(GoMsiExecutable)
		if !temp && (os.IsNotExist(err) || err != nil || fileInfo.IsDir()) {
			log.Println("❌ This Go Version is not downloaded on your machine!")
			if downloadGo() != nil {
				return
			}
		} else if temp {
			if downloadGo() != nil {
				return
			}
			defer os.Remove(GoMsiExecutable) // Remove Temp File
		}

		log.Printf("Installing go%s... ⏳\n", version)
		GoInstaller := gvm.MakeGoInstaller(GoMsiExecutable)
		ok = GoInstaller.InstallAsMSI()
		if !ok {
			log.Printf("Failed to install go%s ❌\n", version)
			os.Remove(GoMsiExecutable) // Remove Corrupted File
			return
		}

		log.Printf("go%s Installed Successfully ✅\n", version)
	},
}

var temp bool

func init() {
	rootCmd.AddCommand(switchCmd)
	switchCmd.Flags().BoolVar(&temp, "no-cache", false, "Whether or not the downloaded Go Version will be cached in disk")
}
