/*
Copyright © 2026 Mike Flot

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/mic615/chill-crate-cli/cmd/buckets"
	"github.com/mic615/chill-crate-cli/cmd/groups"
	"github.com/mic615/chill-crate-cli/cmd/objects"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "chill",
	Short: "Chill Crate object storage, from your terminal",
	Long: `chill is the command-line client for Chill Crate, a simple S3-style
object store. It allows you to create resource groups, manage buckets, and upload/download files from object storage
from your terminal.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(groups.GroupsCmd())
	rootCmd.AddCommand(buckets.BucketsCmd())
	rootCmd.AddCommand(objects.ObjectsCmd())
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().
		StringVar(&cfgFile, "config", "", "config file (default is $HOME/.chill-crate.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.SetConfigFile(filepath.Join(home, ".chill-crate.yaml"))
	}

	viper.AutomaticEnv()
	viper.SetDefault("api_url", "http://localhost:8081")

	if err := viper.SafeWriteConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Writing config file:", viper.ConfigFileUsed())
	}
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
