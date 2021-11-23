package cmd

import (
	"github.com/spf13/cobra"
	"url_shortner/internal/adapters/driving/handleHttp"
)

var (
	rootCmd = &cobra.Command{
		Use:   "Shortner",
		Short: "shorting urls",
		Long:  "shorting long urls",
		Run: func(cmd *cobra.Command, args []string) {
			handleHttp.StartServer()
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}
