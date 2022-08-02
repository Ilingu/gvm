/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"gvm-windows/cmd/utils"
	"gvm-windows/gvm"
	"log"

	"github.com/spf13/cobra"
)

// dlCmd represents the dl command
var dlCmd = &cobra.Command{
	Use:   "dl",
	Short: "dl command downloads the specified Go Version",
	Long:  "Example: `gvm dl 1.18.4`",
	Run: func(cmd *cobra.Command, args []string) {
		if !utils.IsArgsValids(args) {
			cmd.Help()
			return
		}
		version := args[0]

		// Downloading go
		log.Printf("Downloading go%s... ⏳\n", version)
		Godl := gvm.MakeGoDownloader(version)
		dlPath, ok := Godl.DownloadSource()
		if !ok {
			log.Printf("Failed to download go%s ❌\n", version)
			return
		}
		log.Printf("go%s Downloaded Successfully ✅: %s\n", version, dlPath)

		log.Printf("Bundling go%s... ⏳\n", version)
		GoInstaller := gvm.MakeGoInstaller(dlPath, version)
		ok = GoInstaller.InstallAsSource()
		if !ok {
			log.Printf("Failed to bundle go%s ❌\n", version)
			return
		}

		log.Printf("You can now execute `gvm use %s` to enable this version\n", version)
	},
}

func init() {
	rootCmd.AddCommand(dlCmd)
	// goroot := *dlCmd.Flags().String("goroot", `C:\Program Files\Go`, "Custom GOROOT path, default: C:\\Program Files\\Go ")
}
