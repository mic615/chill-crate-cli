/*
Copyright © 2026 Mike Flot
*/
package groups

import (
	"github.com/spf13/cobra"
)

func GroupsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "groups",
		Short: "Manage your groups",
		Long:  ``,
	}

	cmd.AddCommand(createCmd, listCmd, useCmd, currentCmd)
	return cmd
}
