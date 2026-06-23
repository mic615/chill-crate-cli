package groups

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// currentCmd represents the current command
var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Show the active group",
	Long:  `Show the active group`,
	RunE: func(cmd *cobra.Command, args []string) error {
		groupID := viper.GetString("current_group_ID")
		if groupID == "" {
			return fmt.Errorf("no group selected — run 'chill groups use' first")
		}
		fmt.Printf(
			"You are currently using the group : %s \n",
			viper.GetString("current_group_name"),
		)
		return nil
	},
}
