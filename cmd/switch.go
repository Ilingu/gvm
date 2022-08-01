/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"gvm-windows/cmd/utils"
	"gvm-windows/gvm"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Let you switch of go version.",
	Long:  "Let you switch of go version. It unistalls the current Go version, then it downloads the msi executable from Go Official Website and finally it installs the newly downloaded Go version.",
	Run: func(cmd *cobra.Command, args []string) {
		if !utils.IsArgsValids(args) {
			cmd.Help()
			return
		}
		version := args[0]

		if strings.Contains(runtime.Version(), version) {
			log.Printf("You are already on go%s ❌\n", version)
			return
		}

		// Downloading go
		log.Printf("Downloading go%s... ⏳\n", version)
		Godl := gvm.MakeGoDownloader(version)
		dlPath, ok := Godl.DownloadMSI()
		if !ok {
			log.Printf("Failed to download go%s ❌\n", version)
			return
		}
		defer os.Remove(dlPath) // Delete the temp file containing the go executable
		log.Printf("go%s Downloaded Successfully ✅: %s\n", version, dlPath)

		log.Printf("Installing go%s... ⏳\n", version)
		GoInstaller := gvm.MakeGoInstaller(dlPath)
		ok = GoInstaller.InstallAsMSI()
		if !ok {
			log.Printf("Failed to install go%s ❌\n", version)
			return
		}
		log.Printf("go%s Installed Successfully ✅\n", version)
	},
}

func init() {
	rootCmd.AddCommand(switchCmd)
}
