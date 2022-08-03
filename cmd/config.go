/*
Copyright © 2022 BitsOfAByte

*/
package cmd

import (
	"BitsOfAByte/proto/shared"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Change the configuration of Proto",
}

var showConfCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the current configuration",
	Run: func(cmd *cobra.Command, args []string) {
		// Run WriteConfig to make sure a config file exists.
		err := viper.WriteConfig()
		shared.Check(err)

		file, err := os.Open(viper.ConfigFileUsed())
		shared.Check(err)

		defer file.Close()

		config, err := ioutil.ReadAll(file)
		shared.Check(err)

		fmt.Println(string(config))
		fmt.Println("Located at: " + viper.ConfigFileUsed())
	},
}

// verboseCmd represents the verbose command
var verboseCmd = &cobra.Command{
	Use:       "verbose <bool>",
	Short:     "Toggle verbose mode",
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"true", "false"},
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("cli.verbose", args[0])
		viper.WriteConfig()
	},
}

// tempCmd represents the temp command
var tempCmd = &cobra.Command{
	Use:     "temp <dir>",
	Short:   "Change the temporary storage location",
	Example: "proto config temp /tmp/proto/",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Temp location is now: ", shared.UsePath(args[0], true))
		viper.Set("storage.tmp", shared.UsePath(args[0], true))
		viper.WriteConfig()
	},
}

// checksumCmd represents the checksum command
var checksumCmd = &cobra.Command{
	Use:       "force-sum <bool>",
	Short:     "Enable or disable mandatory checksum passing for all downloads",
	Example:   "proto config force-sum true",
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"true", "false"},
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Force sum has been set to", args[0])
		viper.Set("app.forcechecksum", args[0])
		viper.WriteConfig()
	},
}

// installDirCmd represents the install command
var installDirCmd = &cobra.Command{
	Use:     "install <dir>",
	Short:   "Change the proton install directory",
	Example: "proto config install ~/.steam/root/compatibilitytools.d/",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Install location is now: ", shared.UsePath(args[0], true))
		viper.Set("storage.install", shared.UsePath(args[0], true))
		viper.WriteConfig()
	},
}

// resetCmd represents the refresh command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the configuration to default",
	Long:  "Reset the configuration to default, useful for when a major update occurs",
	Run: func(cmd *cobra.Command, args []string) {
		os.Remove(viper.ConfigFileUsed())
		fmt.Println("Configuration has been reset to default")
	},
}

// sourceCmd represents the source command
var sourcesCmd = &cobra.Command{
	Use:   "sources <cmd>",
	Short: "Modify the sources list",
	Args:  cobra.ExactArgs(1),
}

// addSourceCmd represents the add command
var addSourceCmd = &cobra.Command{
	Use:     "add <owner/repo>",
	Short:   "Add a source to the list",
	Example: "proto config sources add <url>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		for _, v := range viper.GetStringSlice("app.sources") {
			if v == args[0] {
				fmt.Println("Source already exists")
				return
			}
		}

		viper.Set("app.sources", append(viper.GetStringSlice("app.sources"), args[0]))
		viper.WriteConfig()
	},
}

// removeSourceCmd represents the remove command
var removeSourceCmd = &cobra.Command{
	Use:     "remove <owner/repo>",
	Short:   "Remove a source from the list",
	Example: "proto config sources remove <url>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// find the source in the list
		var sources = viper.GetStringSlice("app.sources")

		for i, source := range sources {
			if source == args[0] {
				sources = append(sources[:i], sources[i+1:]...)
				break
			}
		}

		viper.Set("app.sources", sources)
		viper.WriteConfig()
	},
}

var listSourcesCmd = &cobra.Command{
	Use:   "list",
	Short: "List all sources",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Currently configured sources:")
		for _, v := range viper.GetStringSlice("app.sources") {
			fmt.Println("- " + v)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.AddCommand(tempCmd)
	configCmd.AddCommand(installDirCmd)
	configCmd.AddCommand(checksumCmd)
	configCmd.AddCommand(verboseCmd)
	configCmd.AddCommand(showConfCmd)
	configCmd.AddCommand(sourcesCmd)
	configCmd.AddCommand(resetCmd)

	sourcesCmd.AddCommand(addSourceCmd)
	sourcesCmd.AddCommand(removeSourceCmd)
	sourcesCmd.AddCommand(listSourcesCmd)
}
