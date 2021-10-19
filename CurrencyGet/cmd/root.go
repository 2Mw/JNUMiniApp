package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:              "Hugo",
	Short:            "Get for currency",
	Long:             "Currency get illustration of commands in AirJ inner network",
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
	cobra.MousetrapHelpText = ""
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Currency",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Currency Get Release 1.0 2021/10/19")
	},
}
