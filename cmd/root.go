/*
Copyright © 2022 Ilingu <ilingu@protonmail.com>
*/

package cmd

import (
	appos "gvm/app_os"
	"gvm/cmd/manager"
	"gvm/console"
	"gvm/utils"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gvm",
	Short: "Simple Go Version Manager for Windows and linux.",
	Long: `A Go Version Manager written in Go that automate the task of switching Go version in Windows.
It downloads the msi file on the official site and execute it.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if !appos.IsSupported() {
		console.Error("You are not on a suported OS ❌")
		return
	}
	if runtime.GOOS == "linux" && !utils.IsExecutedAsRoot() {
		console.Error("❌ You are not root, use `sudo !!`")
		return
	}

	appErr := rootCmd.Execute()
	if appErr != nil {
		console.Error("Fatal Error: cannot invoke app.")
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(manager.ManagerCmds)
}
