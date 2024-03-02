package cmd

import (
	"internal/statuslog"
	"log"

	"github.com/spf13/cobra"
)

var NewStatusCmd = &cobra.Command{
	Use:   "new <staus>",
	Short: "Create a new (one line) status",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		success, err := statuslog.NewStatus(args[0])
		if err != nil {
			log.Fatalln(err)
		} else if !success {
			log.Fatalln("Failed to post status")
		}
	},
}
