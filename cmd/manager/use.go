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
	"strings"

	"github.com/spf13/cobra"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "➡️ Switch between multiples version of Go",
	Long:  `➡️ Switch between multiples version of Go. If the specified Go Version is not downloaded the process exit.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !utils.IsArgsValids(args) {
			cmd.Help()
			return
		}

		targetVersion := args[0]
		UserGoVersion, _ := utils.GetUserGoVersion()

		if targetVersion == "latest" {
			latestVersion, ok := utils.GetLatestGoVersion()
			if !ok {
				log.Println("❌ Cannot fetch latest Go Version. Check your internet connection!")
				return
			}
			targetVersion = latestVersion
		}

		// Check Version
		if strings.Contains(UserGoVersion, targetVersion) {
			log.Printf("You are already on go%s ❌\n", targetVersion)
			return
		}

		appFolder, err := gvmUtils.GenerateAppDataPath()
		if err != nil {
			return
		}

		GoMsiExecutable := appFolder + fmt.Sprintf("\\go%s.msi", targetVersion)
		fileInfo, err := os.Stat(GoMsiExecutable)
		if os.IsNotExist(err) || err != nil || fileInfo.IsDir() {
			log.Printf("❌ This Go Version is not downloaded on your machine!\nType: `gvm manager dl %s` to download this version", targetVersion)
			return
		}

		log.Printf("Switching to go%s... ⏳\n", targetVersion)
		GoInstaller := gvm.MakeGoInstaller(GoMsiExecutable)
		ok := GoInstaller.InstallAsMSI()
		if !ok {
			log.Printf("Failed to switch to go%s ❌\n", targetVersion)
			return
		}

		log.Printf("Switched to go%s Successfully ✅\n", targetVersion)
	},
}

func init() {
}
