package cmd

import (
	"JNUMiniApp/AirJLogin/service"
	"fmt"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get current account information",
	Long:  "Get the information of current network and the account has login.",
	Run: func(cmd *cobra.Command, args []string) {
		ip, _ := service.GetMyIP()
		mac, _ := service.GetMyMAC(ip)
		if len(ip) == 0 {
			ip = "Get Failed."
		}

		if len(mac) == 0 {
			mac = "Get Failed."
		}

		user, _ := service.GetCurrentUser(ip, mac)

		fmt.Printf("Your information:\n\n")
		fmt.Printf("\tIP:\t\t%v\n", ip)
		fmt.Printf("\tMAC:\t\t%v\n", mac)
		fmt.Printf("\tUser:\t\t%v\n", user)
		fmt.Printf("\n\n")
	},
}
