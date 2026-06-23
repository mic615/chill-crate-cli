/*
Copyright © 2026 Mike Flot
*/

package objects

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/mic615/chill-crate-cli/internal/client"
)

var uploadCmd = &cobra.Command{
	Use:   "upload <bucketname> <filePath>",
	Short: "Upload an object to a bucket",
	Long:  `Upload an object to a bucket`,
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		bucketName, fileName := args[0], args[1]
		c := client.New()
		groupID := viper.GetString("current_group_ID")
		if groupID == "" {
			return fmt.Errorf("no group selected — run 'chill groups use' first")
		}
		bucket, err := c.GetBucketByName(bucketName, groupID)
		if err != nil {
			return fmt.Errorf("finding the bucket %w", err)
		}
		// TODO validation check filename etc
		file, err := os.Open(fileName)
		if err != nil {
			return err
		}
		defer file.Close()
		info, err := file.Stat()
		if err != nil {
			return err
		}
		bar := progressbar.DefaultBytes(info.Size(), "uploading")
		fileReader := io.TeeReader(file, bar)
		object, err := c.UploadObject(bucket.ID, filepath.Base(fileName), fileReader, info.Size())
		if err != nil {
			return fmt.Errorf("uploading file %s: %w", fileName, err)
		}
		fmt.Printf("object name: %s created \n", object.FileName)
		return nil
	},
}
