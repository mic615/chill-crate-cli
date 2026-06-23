package objects

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/mic615/chill-crate-cli/internal/client"
)

var downloadCmd = &cobra.Command{
	Use:   "download  <bucketname> <filename> <destination>",
	Short: "Download an object from a bucket",
	Long:  "Download an object from a bucket",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		bucketName, fileName, dest := args[0], args[1], args[2]
		c := client.New()
		groupID := viper.GetString("current_group_ID")
		bucket, err := c.GetBucketByName(bucketName, groupID)
		if err != nil {
			return fmt.Errorf("finding the bucket %w", err)
		}
		body, err := c.DownloadObject(bucket.ID, fileName)
		if err != nil {
			return fmt.Errorf("downloading this object: %w", err)
		}
		defer body.Close()

		out, err := os.Create(dest)
		if err != nil {
			return fmt.Errorf("creating %s: %w", dest, err)
		}
		defer out.Close()
		if _, err := io.Copy(out, body); err != nil {
			return fmt.Errorf("writing %s: %w", dest, err)
		}
		fmt.Printf("downloaded %s -> %s\n", fileName, dest)
		return nil
	},
}
