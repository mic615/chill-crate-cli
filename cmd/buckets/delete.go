package buckets

import (
	"errors"
	"fmt"
	"io"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/mic615/chill-crate-cli/internal/client"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <bucketname>",
	Short: "Delete a bucket",
	Long:  `Delete a bucket.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		force, _ := cmd.Flags().GetBool("force")
		c := client.New()
		groupID := viper.GetString("current_group_id")
		if groupID == "" {
			return fmt.Errorf("no group selected — run 'chill groups use' first")
		}
		bucket, err := c.GetBucketByName(args[0], groupID)
		if err != nil {
			return fmt.Errorf("finding the bucket: %w", err)
		}
		deleteLabel := fmt.Sprintf("Are you sure you want to delete bucket %s", args[0])
		if force {
			deleteLabel = fmt.Sprintf(
				"Are you sure you want to delete bucket %s and all of its contents",
				args[0],
			)
		}
		prompt := promptui.Prompt{
			Label:     deleteLabel,
			IsConfirm: true,
			Default:   "n",
			Stdin:     io.NopCloser(cmd.InOrStdin()),
		}
		_, err = prompt.Run()
		if err != nil {
			if errors.Is(err, promptui.ErrAbort) {
				fmt.Fprintln(cmd.OutOrStdout(), "Delete canceled.")
				return nil // user declined
			}
			return err
		}
		if err := c.DeleteBucket(bucket.ID, force); err != nil {
			return err
		}

		fmt.Fprintf(cmd.OutOrStdout(), "%s successfully deleted\n", args[0])
		return nil
	},
}

func init() {
	deleteCmd.Flags().
		BoolP("force", "f", false, "force deletes all objects in a bucket before deleting the bucket.")
}
