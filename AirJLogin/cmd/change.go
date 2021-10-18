package cmd

import (
	"JNUMiniApp/AirJLogin/service"
	"github.com/spf13/cobra"
	"log"
	"time"
)

var changeCmd = &cobra.Command{
	Use:   "change",
	Short: "Change AirJ account",
	Long:  "Change a new account of airJ, then login with this account, if this account exists, then change to this.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(account) != 10 {
			log.Println("Your account length is invalid")
			return
		}

		if len(password) >= 6 || len(password) == 0 {
			//log.Println(account, password)
			b := service.AddLoginData(account, &password)
			if b {
				log.Printf("Change successfully. Start to login %v...\n", account)
				service.Login(account, password)
			}else{
				log.Println("Change failed")
			}
		}else {
			log.Println("Password length invalid.")
		}

		time.Sleep(time.Second * 2)
	},
}

func init() {
	changeCmd.Flags().StringVarP(&account, "account", "a", "", "your account (required)")
	changeCmd.Flags().StringVarP(&password, "password", "p", "", "your password")
	_ = changeCmd.MarkFlagRequired("account")
}
