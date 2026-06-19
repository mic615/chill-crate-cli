/*
Copyright © 2026 Mike Flot
*/

package buckets

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/mic615/chill-crate-cli/internal/client"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your buckets",
	Long:  `List all buckets in your current group.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c := client.New()
		groupID := viper.GetString("current_group_ID")
		buckets, err := c.ListBuckets(groupID)
		if err != nil {
			return fmt.Errorf("getting buckets: %w", err)
		}
		if len(buckets) == 0 {
			fmt.Printf("No buckets yet — create one with chill groups bucket <name>")
			return nil
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		for _, b := range buckets {
			fmt.Fprintf(w, "%s\t%s\n", b.Name, b.ID)
		}
		return w.Flush()
	},
}
