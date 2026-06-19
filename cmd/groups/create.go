/*
Copyright © 2026 Mike Flot
*/

package groups

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mic615/chill-crate-cli/internal/client"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "Create a new group",
	Long:  `Create a new group with the given name. You are automatically added as a member.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c := client.New()
		group, err := c.CreateGroup(args[0])
		if err != nil {
			return fmt.Errorf("creating group %s: %w", args[0], err)
		}
		fmt.Printf("group name: %s created \n", group.Name)
		return nil
	},
}
