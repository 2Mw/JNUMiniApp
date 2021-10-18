package cmd

import (
	"JNUMiniApp/AirJLogin/service"
	"fmt"
	"github.com/spf13/cobra"
)

type void struct{}

var zero void

var listCmd = &cobra.Command{
	Use: "list",
	Short: "list all user and his currency",
	Long: "List all users and details of their currency",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := service.ReadContent()
		if err != nil {
			return
		}

		set := make(map[string]void)
		set[data.Acc] = zero
		for _, item := range data.Alternatives {
			set[item.Acc] = zero
		}
		// Get
		for k := range set {
			txt := ""
			if k == data.Acc{
				txt = "*"
			}
			fmt.Printf("User: %v \tCurrecy: %v %v\n", k, service.GetCurrency(k), txt)
		}
	},
}
