package cmd

import (
	"JNUMiniApp/CurrencyGet/service"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

const (
	version     = "v1.1.220606"
	releaseDate = "2022/06/06"
)

var threads = 0

var rootCmd = &cobra.Command{
	Use:              "CurrencyGet",
	Short:            "Get for currency",
	Long:             "Currency get illustration of commands in AirJ inner network",
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		Attention()
		if threads <= 0 || threads > 500 {
			log.Printf("%v\n", "Threads must lower than 500 and greater than 0")
			return
		}
		//fmt.Println(threads)
		service.StartThreads(threads)
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Currency",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Currency Get Release %v at %v\n", version, releaseDate)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.Flags().IntVarP(&threads, "thread", "t", 80, "Specify work threads, recommend 80.")
	cobra.MousetrapHelpText = ""
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func Attention() {
	fmt.Println("=============================  WARNING  =============================")
	fmt.Println("THIS PROGRAM IS ONLY FOR STUDY USE. PLEASE DON'T SPREAD IT TO")
	fmt.Println("INTERNET, OTHERWISE ALL THE LEGAL CONSEQUENCES ARISING THEREFROM")
	fmt.Println("SHALL BE BORNE BY THE USER. ALL THE CONSEQUENCES HAVE NOTHING WITH")
	fmt.Println("THE DEVELOPER. PLEASE DELETE THIS SOFTWARE WITHIN 24 HOURS.")
	fmt.Println("=============================  WARNING  =============================")
	fmt.Printf("\n※Release Version: %v, release time: %v※\n", version, releaseDate)
	fmt.Println("\n")
}
