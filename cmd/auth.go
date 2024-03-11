package cmd

import (
	"fmt"
	"internal/token"
	"os"
	
	"github.com/spf13/cobra"
)

var AuthCmd = &cobra.Command {
	Use: "auth <cmd>",
	Short: "Setup or view auth",
	Args: cobra.MinimumNArgs(1),
	
	Run: func(cmd *cobra.Command, args [] string) {
		switch args[0] {
		case "status":
			token := token.GetTokenEnv()
			if token == "" {
				fmt.Println("Token not set")
				os.Exit(1)
			}
			// TODO: confirm by using the omglol api
			fmt.Println("Token is set!")
			os.Exit(0)
		case "token":
			fmt.Printf("\"%s\"\n", token.GetTokenEnv())
			os.Exit(0)
		case "set":
			// TODO: check that 2 args are given
			token.SetTokenEnv(args[1])
			os.Exit(0)
			
			/*
			fmt.Println("auth set <TOKEN>")
			os.Exit(1)
			*/
		}
	}}

