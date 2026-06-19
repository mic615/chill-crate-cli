/*
Copyright © 2026 Mike Flot
*/

package objects

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/mic615/chill/internal/client"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list  <bucket>",
	Short: "List all the objects in your bucket",
	Long:  `List all the objects in your bucket.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c := client.New()
		objects, err := c.ListObjects(args[0])

		if err != nil {
			return fmt.Errorf("getting objects: %w", err)
		}
		if len(objects) == 0 {
			fmt.Printf("No objects yet — create one with chill objects upload <name>")
			return nil
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

		fmt.Fprintf(w, "File Name\t version\n")
		for _, o := range objects {
			fmt.Fprintf(w, "%s\t%d\n", o.FileName, o.Version)
		}
		return w.Flush()
	},
}
