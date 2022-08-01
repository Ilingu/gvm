/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"gvm-windows/cmd/utils"

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
	},
}

func init() {
	rootCmd.AddCommand(dlCmd)
	// goroot := *dlCmd.Flags().String("goroot", `C:\Program Files\Go`, "Custom GOROOT path, default: C:\\Program Files\\Go ")
}
