package buckets

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"github.com/mic615/chill-crate-cli/internal/client"
)

var force bool

var deleteCmd = &cobra.Command{
	Use:   "delete <bucketname>",
	Short: "Delete a bucket",
	Long:  `Delete a bucket.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c := client.New()
		deleteLabel := fmt.Sprintf("Are you sure you want to delete bucket %s", args[0])
		if force {
			deleteLabel = fmt.Sprintf(
				"Are you sure you want to delete bucket %s and all of it's contents",
				args[0],
			)
		}
		prompt := promptui.Prompt{
			Label:     deleteLabel,
			IsConfirm: true,
		}
		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

		if strings.ToLower(result) != "y" {
			return nil
		}

		return c.DeleteBucket(args[0], force)
	},
}

func init() {
	deleteCmd.Flags().BoolVarP(&force, "force", "f", false, "")
}
