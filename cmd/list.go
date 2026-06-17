/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/mic615/chill/internal/client"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your groups",
	Long:  `List all groups you are a member of.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c := client.New()
		groups, err := c.ListGroups()

		if err != nil {
			return fmt.Errorf("getting groups: %w", err)
		}
		for _, group := range groups {

			fmt.Printf("group name: %s\n", group.Name)
		}
		return nil
	},
}

func init() {
	groupsCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
