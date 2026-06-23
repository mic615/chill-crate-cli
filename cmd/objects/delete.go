package objects

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/mic615/chill-crate-cli/internal/client"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <bucketname> <filename> ",
	Short: "Delete an file from a bucket",
	Long:  "Delete an file from a bucket",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		bucketName, fileName := args[0], args[1]
		c := client.New()
		groupID := viper.GetString("current_group_ID")
		bucket, err := c.GetBucketByName(bucketName, groupID)
		if err != nil {
			return fmt.Errorf("finding the bucket %w", err)
		}
		if err := c.DeleteObject(bucket.ID, fileName); err != nil {
			return fmt.Errorf("deleting this file %w", err)
		}
		fmt.Printf("deleted %s \n", fileName)
		return nil
	},
}
