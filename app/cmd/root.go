package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "event",
	Short: "event is a simple event-driven manager",
	Long:  `此处省略 80000 字介绍`,
}
