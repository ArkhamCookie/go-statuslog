package cmd

import (
	"fmt"
	"internal/statuslog"
	"os"

	"github.com/spf13/cobra"
)

var BioViewCmd = &cobra.Command{
	Use: "bio <address>",
	Short: "View an address' bio",
	Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		bio, err := statuslog.BioGet(args[0])
		if err != nil {
			if bio.Request.StatusCode == 404 {
				fmt.Println(bio.Response.Message)
				os.Exit(0)
			}

			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(bio.Response.Bio)
	},
}
