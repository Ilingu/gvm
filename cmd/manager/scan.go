/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package manager

import (
	"gvm-windows/gvm"
	"log"
	"time"

	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "➡️ Scan and delete old downloaded Go MSI file",
	Long:  `➡️ Scan and delete old downloaded Go MSI file based on the file creation date.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Scanning... ⏳")
		filesDeleted, err := gvm.ScanAndDelete(limitDate)
		if err != nil {
			log.Println("❌ Scanning failed!")
			return
		}

		log.Printf("✅ Scanning completed! %d files deleted.", filesDeleted)
	},
}

var limitDate int64

const SIX_MONTH_IN_MS = 15778800000 // 1000 * 60 * 60 * 24 * 182.625 --> There is exactly 182.625 days in 6 month (365.25/2)

func init() {
	SixMonthAgo := time.Now().UnixMilli() - SIX_MONTH_IN_MS
	scanCmd.Flags().Int64VarP(&limitDate, "date-limit", "l", SixMonthAgo, "All the Go Version downloaded before this date will be deleted. I expect a timestamp of this date in MILLISECOND. Default= ~6 month ago.")
}
