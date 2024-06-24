package cmd

import (
	"fmt"
	"internal/statuslog"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

var LatestCmd = &cobra.Command{
	Use: "latest <address>",
	Short: "View an address' latest status",
	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		// Get address from args
		address := args[0]

		// Retreive statuses
		result, err := statuslog.ListGet(address)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Format output to print only the latest status
		output := fmt.Sprintf(
			"%s %s %s\n",

			result.Response.Statuses[0].Emoji,
			result.Response.Statuses[0].Content,
			result.Response.Statuses[0].RelativeTime,
		)

		// Use glamour to render output
		out, err := glamour.Render(output, "dark")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Print rendered output
		fmt.Println(out)
	},
}
