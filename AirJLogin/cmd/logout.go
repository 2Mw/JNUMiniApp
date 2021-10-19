package cmd

import (
	"JNUMiniApp/AirJLogin/service"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout current account",
	Long:  "Logout current account in inner network",
	Run: func(cmd *cobra.Command, args []string) {
		service.Logout()
	},
}
