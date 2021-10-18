package cmd

import (
	"JNUMiniApp/AirJLogin/service"
	"github.com/spf13/cobra"
	"log"
)

var delCmd = &cobra.Command{
	Use: "del",
	Short: "delete a account info",
	Long: "delete a specified account info",
	Run: func(cmd *cobra.Command, args []string) {
		if len(account) == 10 {
			b := service.DelAccount(account)
			if b {
				log.Printf("%v\n", "delete successfully")
			}else {
				log.Println("delete failed")
			}
		}else{
			log.Println("Account is invalid.")
		}
	},
}

func init() {
	delCmd.Flags().StringVarP(&account, "account", "a", "", "Specify an account")
	_ = delCmd.MarkFlagRequired("account")
}
