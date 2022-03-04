package cmd

import (
	"github.com/flc1125/event/cron"
	"github.com/spf13/cobra"
)

var cronCmd = &cobra.Command{
	Use:   "cron",
	Short: "Run cron",
	Run: func(cmd *cobra.Command, args []string) {
		cron.Run()
	},
}

func init() {
	rootCmd.AddCommand(cronCmd)
}
