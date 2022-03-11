package cmd

import (
	"github.com/flc1125/event/app/http"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run server",
	Run: func(cmd *cobra.Command, args []string) {
		http.Start()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
