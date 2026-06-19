/*
Copyright © 2026 Mike Flot
*/

package buckets

import (
	"fmt"

	"github.com/mic615/chill/internal/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "Create a new bucket",
	Long:  `Create a new bucket in your current group with the given name.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c := client.New()
		groupID := viper.GetString("current_group_ID")
		bucket, err := c.CreateBucket(args[0], groupID)
		if err != nil {
			return fmt.Errorf("creating bucket %s: %w", args[0], err)
		}
		fmt.Printf("bucket name: %s created \n", bucket.Name)
		return nil
	},
}
