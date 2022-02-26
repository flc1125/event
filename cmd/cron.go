package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var cronCmd = &cobra.Command{
	Use:   "cron",
	Short: "Run cron",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("asdfasfd")
	},
}

func init() {
	rootCmd.AddCommand(cronCmd)
}
