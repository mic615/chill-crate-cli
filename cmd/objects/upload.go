/*
Copyright © 2026 Mike Flot
*/

package objects

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/mic615/chill/internal/client"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:   "upload <bucket> <filePath>",
	Short: "Upload an object to a bucket",
	Long:  `Upload an object to a bucket`,
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		c := client.New()
		// TODO validation check filename etc
		file, err := os.Open(args[1])
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
		object, err := c.UploadObject(args[0], filepath.Base(args[1]), fileReader, info.Size())
		if err != nil {
			return fmt.Errorf("uploading file %s: %w", args[1], err)
		}
		fmt.Printf("object name: %s created \n", object.FileName)
		return nil
	},
}
