package objects

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mic615/chill-crate-cli/internal/client"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <bucket> <filename> ",
	Short: "Delete an file from a bucket",
	Long:  "Delete an file from a bucket",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		bucket, fileName := args[0], args[1]
		c := client.New()
		if err := c.DeleteObject(bucket, fileName); err != nil {
			return fmt.Errorf("deleting this file %w", err)
		}
		fmt.Printf("deleted %s \n", fileName)
		return nil
	},
}
