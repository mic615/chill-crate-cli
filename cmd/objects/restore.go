package objects

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/mic615/chill-crate-cli/internal/client"
)

var restoreCmd = &cobra.Command{
	Use:   "restore <bucketname> <filename> ",
	Short: "Restore a deleted file",
	Long:  "Restore a deleted file to a bucket",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		bucketName, fileName := args[0], args[1]
		c := client.New()
		groupID := viper.GetString("current_group_ID")
		bucket, err := c.GetBucketByName(bucketName, groupID)
		if err != nil {
			return fmt.Errorf("finding the bucket %w", err)
		}
		if err := c.RestoreObject(bucket.ID, fileName); err != nil {
			return fmt.Errorf("restoring this file %w", err)
		}
		fmt.Printf("restored %s \n", fileName)
		return nil
	},
}
