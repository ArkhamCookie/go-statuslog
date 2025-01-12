package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sl [flags]",
	Short: "Statuslog is a commandline interface for omg.lol's statuslog feature",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("statuslog")
	},
}

func init() {
	rootCmd.AddCommand(AuthCmd)
	rootCmd.AddCommand(BioViewCmd)
	rootCmd.AddCommand(LatestCmd)
	rootCmd.AddCommand(ListCmd)
	rootCmd.AddCommand(NewStatusCmd)
	rootCmd.AddCommand(StatusViewCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
