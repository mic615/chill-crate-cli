/*
Copyright © 2026 Mike Flot
*/

package objects

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
	Use:   "list  <bucketname>",
	Short: "List all the objects in your bucket",
	Long:  `List all the objects in your bucket.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c := client.New()
		groupID := viper.GetString("current_group_ID")
		if groupID == "" {
			return fmt.Errorf("no group selected — run 'chill groups use' first")
		}
		bucket, err := c.GetBucketByName(args[0], groupID)
		if err != nil {
			return fmt.Errorf("finding the bucket %w", err)
		}
		objects, err := c.ListObjects(bucket.ID)
		if err != nil {
			return fmt.Errorf("getting objects: %w", err)
		}
		if len(objects) == 0 {
			fmt.Printf("No objects yet — create one with chill objects upload <name>")
			return nil
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

		fmt.Fprintf(w, "File Name\tversion\n")
		for _, o := range objects {
			fmt.Fprintf(w, "%s\t%d\n", o.FileName, o.Version)
		}
		return w.Flush()
	},
}
