/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// groupsCmd represents the groups command
var groupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "Manage your groups",
	Long:  ``,
}

func init() {
	rootCmd.AddCommand(groupsCmd)
}
