/*
Copyright © 2026 Mike Flot
*/

package groups

import (
	"fmt"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/mic615/chill-crate-cli/internal/client"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your groups",
	Long:  `List all groups you are a member of.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c := client.New()
		groups, err := c.ListGroups()
		if err != nil {
			return fmt.Errorf("getting groups: %w", err)
		}
		if len(groups) == 0 {
			fmt.Fprintf(
				cmd.OutOrStdout(),
				"No groups yet — create one with chill groups create <name>\n",
			)
			return nil
		}
		current := viper.GetString("current_group_ID")

		w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 3, ' ', 0)
		for _, g := range groups {
			marker := " "
			if g.ID == current {
				marker = "*"
			}
			fmt.Fprintf(w, "%s\t%s\t%s\n", marker, g.Name, g.ID)
		}
		return w.Flush()
	},
}
