package cmd

import (
	"fmt"
	"internal/statuslog"
	"os"

	"github.com/spf13/cobra"
)

var LatestCmd = &cobra.Command{
	Use: "",
	Short: "",
	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		address := args[0]

		result, err := statuslog.LatestGet(address)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Sprintln(result)
	},
}
