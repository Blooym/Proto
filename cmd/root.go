/*
Copyright © 2022 BitsOfAByte

GPLv3 License, see the LICENSE file for more information.
*/
package cmd

import (
	"BitsOfAByte/proto/config"
	"BitsOfAByte/proto/core"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:     "proto",
	Short:   "Install and manage custom runners with ease ",
	Version: core.Version,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Register persistent flags
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose output")
	rootCmd.PersistentFlags().BoolP("yes", "y", false, "Skip all confirmation prompts")
	rootCmd.PersistentFlags().StringP("dir", "d", "", "The directory to operate in")

	// Register flags to config
	viper.BindPFlag("cli.verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

// Initialize proto configuration file
func initConfig() {
	config.SetDefaults()
}
