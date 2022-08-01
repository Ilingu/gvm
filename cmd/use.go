/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"gvm-windows/cmd/utils"

	"github.com/spf13/cobra"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "'use' will switch from a version A to B.",
	Long:  `It will switch from a version A to B. If version B isn't installated, it will stop.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !utils.IsArgsValids(args) {
			cmd.Help()
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(useCmd)

	// goroot := os.Getenv("GOROOT")
}
