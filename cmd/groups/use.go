/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/

package groups

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/mic615/chill/internal/client"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Set the active group for later commands",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		c := client.New()
		groups, err := c.ListGroups()
		if err != nil {
			return fmt.Errorf("getting groups: %w", err)
		}
		if len(groups) == 0 {
			fmt.Printf("No groups yet — create one with chill groups create <name>")
			return nil
		}
		// current := viper.GetString("current_group")
		groupMap := make(map[string]string)
		var groupNames []string
		for _, g := range groups {
			groupMap[g.Name] = g.ID
			groupNames = append(groupNames, g.Name)
		}

		prompt := promptui.Select{
			Label: "Select a group",
			Items: groupNames,
		}
		_, result, err := prompt.Run()
		if err != nil {
			return fmt.Errorf("prompt failed %w", err)
		}
		viper.Set("current_group_name", result)
		viper.Set("current_group_ID", groupMap[result])
		if err := viper.WriteConfig(); err != nil {
			return fmt.Errorf("updating group: %w", err)
		}
		fmt.Printf("You choose %q, %q\n", result, groupMap[result])
		return nil
	},
}
