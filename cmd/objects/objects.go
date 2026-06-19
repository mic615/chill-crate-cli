/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package objects

import (
	"github.com/spf13/cobra"
)

func ObjectsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "objects",
		Short: "Manage your objects",
		Long:  ``,
	}

	cmd.AddCommand(uploadCmd, downloadCmd, listCmd)
	return cmd
}
