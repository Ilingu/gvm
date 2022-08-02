/*
Copyright © 2022 Ilingu <ilingu@protonmail.com>
*/

package manager

import (
	"fmt"
	"gvm-windows/cmd/utils"
	"gvm-windows/gvm"
	gvmUtils "gvm-windows/gvm/utils"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "➡️ Switch between multiples version of Go",
	Long: `➡️ Switch between multiples version of Go.
	Download source version if not already downloaded (@see the 'dl' command). It then change your GOROOT to the downloaded package, bundles/installs it.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !utils.IsArgsValids(args) {
			cmd.Help()
			return
		}

		version := args[0]
		appFolder, err := gvmUtils.GenerateAppDataPath()
		if err != nil {
			return
		}

		GoMsiExecutable := appFolder + fmt.Sprintf("\\go%s.msi", version)

		fileInfo, err := os.Stat(GoMsiExecutable)
		if os.IsNotExist(err) || err != nil || fileInfo.IsDir() {
			log.Printf("❌ This Go Version is not downloaded on your machine!\nType: `gvm manager dl %s` to download this version", version)
			return
		}

		log.Printf("Switching to go%s... ⏳\n", version)
		GoInstaller := gvm.MakeGoInstaller(GoMsiExecutable)
		ok := GoInstaller.InstallAsMSI()
		if !ok {
			log.Printf("Failed to switch to go%s ❌\n", version)
			return
		}

		log.Printf("Switched to go%s Successfully ✅\n", version)
	},
}

func init() {
}
