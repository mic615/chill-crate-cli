package objects

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mic615/chill-crate-cli/internal/client"
)

var restoreCmd = &cobra.Command{
	Use:   "restore <bucket> <filename> ",
	Short: "Restore a deleted file",
	Long:  "Restore a deleted file to a bucket",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		bucket, fileName := args[0], args[1]
		c := client.New()
		if err := c.RestoreObject(bucket, fileName); err != nil {
			return fmt.Errorf("restoring this file %w", err)
		}
		fmt.Printf("restored %s \n", fileName)
		return nil
	},
}
