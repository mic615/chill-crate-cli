/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package buckets

import (
	"github.com/spf13/cobra"
)

func BucketsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buckets",
		Short: "Manage your buckets",
		Long:  ``,
	}

	cmd.AddCommand(createCmd, listCmd, deleteCmd)
	return cmd
}
