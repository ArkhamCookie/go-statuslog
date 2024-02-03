package cmd

import (
	"fmt"
	"internal/statuslog"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

var StatusViewCmd = &cobra.Command{
	Use: "view <address> <status>",
	Short: "View a status",
	Args: cobra.MinimumNArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		result, err := statuslog.StatusGet(args[0], args[1])
		if err != nil {
			if result.Request.StatusCode == 404 {
				fmt.Println(result.Response.Message)
				os.Exit(0)
			}

			fmt.Println(err)
			os.Exit(1)
		}

		output := fmt.Sprintf(
			"%s %s %s\n",

			result.Response.Status.Emoji,
			result.Response.Status.Content,
			result.Response.Status.RelativeTime,
		)

		out, err := glamour.Render(output, "dark")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(out)
	},
}
