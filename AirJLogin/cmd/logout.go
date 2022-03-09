package cmd

import (
	"JNUMiniApp/AirJLogin/service"
	"github.com/spf13/cobra"
	"log"
)

var ip string
var mac string

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout current account",
	Long:  "Logout current account in inner network.\n ※Attention※: Only unbind mac address can logout completely.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(ip) == 0 {
			myIp, err := service.GetMyIP()
			if err != nil {
				log.Printf("Cant get your inner IP: %v", err)
				return
			}
			ip = myIp
		}
		if len(mac) == 0 {
			myMAC, err := service.GetMyMAC(ip)
			if err != nil {
				log.Printf("Cant get your MAC address: %v", err)
				return
			}
			mac = myMAC
		}
		service.Logout(ip, mac)
	},
}

func init() {
	logoutCmd.Flags().StringVarP(&ip, "ip", "i", "", "The ip you want to unregister.")
	logoutCmd.Flags().StringVarP(&mac, "mac", "m", "", "The mac you want to unregister.")
}
