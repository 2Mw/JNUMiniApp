package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Long:  "Show AirJ Login program version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ArijLogin Version 1.0.211104")
	},
}
