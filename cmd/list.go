package cmd

import (
	"fmt"
	"internal/statuslog"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list <address>",
	Short: "List all of an address' statuses",
	Args:  cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		address := args[0]

		result, err := statuslog.ListGet(address)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		header := fmt.Sprintf("# %s's Statuses\n", address)
		statues := statuslog.ListEach(result)
		output := header + statues

		out, err := glamour.Render(output, "dark")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Print(out)

	},
}
