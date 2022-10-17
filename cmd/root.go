/*
Copyright © 2022 BitsOfAByte

GPLv3 License, see the LICENSE file for more information.
*/
package cmd

import (
	"os"

	"github.com/BitsOfAByte/proto/config"
	"github.com/BitsOfAByte/proto/core"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use:     "proto",
	Short:   "Install and manage custom runners with ease ",
	Version: core.Version,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Register persistent flags
	RootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose output")
	RootCmd.PersistentFlags().BoolP("yes", "y", false, "Skip all confirmation prompts")
	RootCmd.PersistentFlags().StringP("dir", "d", "", "The directory to operate in")

	// Register flags to config
	viper.BindPFlag("cli.verbose", RootCmd.PersistentFlags().Lookup("verbose"))
}

// Initialize proto configuration file
func initConfig() {
	config.SetDefaults()
}
