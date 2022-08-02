/*
Copyright © 2022 Ilingu <ilingu@protonmail.com>
*/

package manager

import (
	"github.com/spf13/cobra"
)

// ManagerCmds represents the go mutliples version manager
var ManagerCmds = &cobra.Command{
	Use:   "manager",
	Short: "➡️ Go Multiple Versions Manager",
	Long:  "➡️ Go Multiple Versions Manager.\nYou can download multiple version of Go (with 'gvm manager dl <go_version>') and switch between them (with 'gvm manager use <go_version>')",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	ManagerCmds.AddCommand(dlCmd)
	ManagerCmds.AddCommand(useCmd)
}
