package cmd

import (
	"JNUMiniApp/AirJLogin/service"
	"github.com/spf13/cobra"
	"log"
	"os"
	"time"
)

var (
	account  string
	password string
)

var rootCmd = &cobra.Command{
	Use:              "",
	Long:             "Illustration of AirJ Login program",
	TraverseChildren: true, // 允许在主命令上使用flag
	Run: func(cmd *cobra.Command, args []string) {
		acc, pass := service.ReadLoginData() // 获取配置文件
		if len(acc) > 0 && len(pass) > 0 {
			service.Login(acc, pass)
		}else{
			log.Println("Account parameter error.")
		}
		time.Sleep(time.Second * 2)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println("Input command format error.")
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd, changeCmd, listCmd, delCmd)
	cobra.MousetrapHelpText = ""
}