package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var AuthCmd = &cobra.Command{
	Use:   "auth <command>",
	Short: "Setup or view auth",
	Long: `Setup or view auth

COMMANDS:
	status:   view auth status
	token:    print current api token
	address:  print current address`,
	Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "status":
			token := os.Getenv("OMGLOL_TOKEN")
			if token == "" {
				fmt.Println("Token not set")
				os.Exit(1)
			}
			// TODO: confirm by using the omglol api
			fmt.Println("Token is set!")
			os.Exit(0)
		case "token":
			fmt.Println(os.Getenv("OMGLOL_TOKEN"))
			os.Exit(0)
		case "address":
			fmt.Println(os.Getenv("OMGLOL_ADDRESS"))
			os.Exit(0)
		default:
			fmt.Printf("Unknown command '%s'\n", args[0])
		}
	}}
