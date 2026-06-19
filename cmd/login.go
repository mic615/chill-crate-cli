/*
Copyright © 2026 Mike Flot
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login <user>",
	Short: "Log in and save your session locally",
	Long: `Log in to Chill Crate and store your session locally so later commands are authenticated.
For now it just records the user you act as; the actual authentication is coming soon.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		user := args[0]
		viper.Set("user", user)
		if err := viper.WriteConfig(); err != nil {
			return fmt.Errorf("saving session: %w", err)
		}
		fmt.Printf("Logged in as %s\n", user)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
