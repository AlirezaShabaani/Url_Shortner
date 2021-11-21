package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "Shortner",
		Short: "shorting urls",
		Long:  "shorting long urls",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print("Helooooooooooooooooo beautiful humans")
		},
	}
)

func Execute() error{
	return rootCmd.Execute()
}
