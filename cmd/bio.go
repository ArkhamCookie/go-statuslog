package cmd

import (
	"fmt"
	"internal/statuslog"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

var BioViewCmd = &cobra.Command{
	Use:   "bio <view|edit> <address|new_bio>",
	Short: "View an address' bio or edit your address' bio",
	Args:  cobra.MinimumNArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "view":
			result, err := statuslog.BioGet(args[1])
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
		case "edit":
			success, err := statuslog.BioEdit(args[1])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			} else if !success {
				fmt.Println("Failed to edit bio")
			}
		}
	},
}
