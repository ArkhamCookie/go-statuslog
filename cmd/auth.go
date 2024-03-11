package cmd

import (
	"fmt"
	"internal/env"
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
	env:      use a file for var's values (defaults to .env)
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
			token := os.Getenv("OMGLOL_TOKEN")
			fmt.Printf("\"%s\"\n", token)
			
			os.Exit(0)
		case "address":
			token := os.Getenv("OMGLOL_ADDRESS")
			fmt.Printf("\"%s\"\n", token)
			
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
		case "env":
				if len(args) == 2 {
					apiKey, err := env.GetEnvValue(args[1], "OMGLOL_TOKEN")
					if err != nil {
						os.Exit(1)
					}
					os.Setenv("OMGLOL_TOKEN", apiKey)
					address, err := env.GetEnvValue(args[1], "OMGLOL_TOKEN")
					if err != nil {
						os.Exit(1)
					}
					os.Setenv("OMGLOL_ADDRESS", address)
					os.Exit(0)
				}
			}
		}
	}}
