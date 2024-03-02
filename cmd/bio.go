package cmd

import (
	"fmt"
	"internal/statuslog"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

var BioViewCmd = &cobra.Command{
	Use:   "bio <address>",
	Short: "View an address' bio",
	Args:  cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		result, err := statuslog.BioGet(args[0])
		if err != nil {
			if result.Request.StatusCode == 404 {
				fmt.Println(result.Response.Message)
				os.Exit(0)
			}

			fmt.Println(err)
			os.Exit(1)
		}

		out, err := glamour.Render(result.Response.Bio, "dark")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Print(out)
	},
}
