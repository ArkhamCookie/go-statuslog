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
	set:      set address or api token
	token:    print current api token
	address:  print current address`,
	Args:  cobra.MinimumNArgs(1),

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
		case "set":
		
			if len(args) < 3 {
				fmt.Println("auth set <address|token> <value>")
				os.Exit(0)
			}
			switch args[1] {
			case "token":
				os.Setenv("OMGLOL_TOKEN", args[2])
				os.Exit(0)
			case "address":
				os.Setenv("OMGLOL_ADDRESS", args[2])
				os.Exit(0)
			default:
				fmt.Println("auth set <address|token> <value>")
				os.Exit(1)
			}
		}
	}}
