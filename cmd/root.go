/*
Copyright © 2022 Ilingu <ilingu@protonmail.com>
*/

package cmd

import (
	"gvm-windows/cmd/manager"
	"log"
	"runtime"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gvm",
	Short: "Simple Go Version Manager for Windows.",
	Long: `A Go Version Manager written in Go that automate the task of switching Go version in Windows.
It downloads the msi file on the official site and execute it.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if runtime.GOOS != "windows" {
		log.Fatal("You are not on windows ❌")
		return
	}

	appErr := rootCmd.Execute()
	if appErr != nil {
		log.Fatal("Fatal Error: cannot invoke app.")
	}
}

func init() {
	rootCmd.AddCommand(manager.ManagerCmds)
}
