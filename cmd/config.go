/*
Copyright © 2022 BitsOfAByte

GPLv3 License, see the LICENSE file for more information.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/BitsOfAByte/proto/core"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Change the configuration of Proto",
	Long:  `Configure how Proto works and behaves by changing its configuration file through safe and easy to use commands.`,
}

var configDirCmd = &cobra.Command{
	Use:   "dir",
	Short: "View the directory where the configuration file is stored",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(viper.ConfigFileUsed())
	},
}

var showConfCmd = &cobra.Command{
	Use:     "show",
	Short:   "Show the current configuration",
	Long:    `Outputs the contents of the configuration file that Proto is using.`,
	Example: `proto config show`,
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		err := viper.WriteConfig()
		core.CheckError(err)

		file, err := os.Open(viper.ConfigFileUsed())
		core.CheckError(err)

		defer file.Close()

		config, err := ioutil.ReadAll(file)
		core.CheckError(err)

		fmt.Println(string(config))
		fmt.Println("Located at: " + viper.ConfigFileUsed())
	},
}

var verboseCmd = &cobra.Command{
	Use:       "verbose <bool>",
	Short:     "Toggle verbose mode",
	Long:      `Toggles verbose mode on or off. When verbose mode is on, Proto will output more information about what it is doing.`,
	Example:   "proto config verbose true",
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"true", "false"},
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("cli.verbose", args[0])
		viper.WriteConfig()
		fmt.Println("Verbose mode is now", args[0])
	},
}

var tempCmd = &cobra.Command{
	Use:     "temp <dir>",
	Short:   "Change the temporary storage location",
	Long:    `Change the location where Proto stores temporary files. Typically you do not want to change this as it is automatically set to the system's temporary directory which is automatically cleaned up for you/`,
	Example: "proto config temp /tmp/proto/",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("storage.tmp", core.UsePath(args[0], true))
		viper.WriteConfig()
		fmt.Println("Temporary file storage location changed to: " + args[0])
	},
}

var forceCmd = &cobra.Command{
	Use:   "force <bool>",
	Short: "Forces installations to go ahead regardless of missing or invalid checksums",
	Long: `Enable or disable forcing installations regardless of missing or invalid checksums.
When disabled, Proto will not allow you to download a file unless it has a valid checksum, if the release does not provide a checksum then it cannot be downloaded.
This is useful for ensuring that the file you downloaded is the same as the one that the author uploaded.
Please note that this does not mean that the file is safe to run if it passes, and you should always trust the source you download from.`,
	Example:   "proto config force-sum true",
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"true", "false"},
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("app.force", args[0])
		viper.WriteConfig()
		fmt.Println("Always force has been set to: " + args[0])
	},
}

var locationsCmd = &cobra.Command{
	Use:   "locations <cmd>",
	Short: "Manage your custom directory mappings",
	Long: `Custom directory mappings allow you to map a directory to a name, and then use that name in place of the directory when downloading files using the --dir flag.
They are very useful for when you have to work with the same directory multiple times and you don't want to constantly re-type the directory name by hand/`,
	Args: cobra.MinimumNArgs(1),
}

var addLocationCmd = &cobra.Command{
	Use:     "add <name> <dir>",
	Short:   "Add a custom location for the --dir flag",
	Example: "proto config locations add steam ~/.steam/root/compatibilitytools.d/",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		// Replace any /home/<username> reference with ~ for portability.
		homeDir, _ := os.UserHomeDir()
		if strings.HasPrefix(args[1], homeDir) {
			args[1] = strings.Replace(args[1], homeDir, "~", 1)
		}

		if strings.Contains(args[0], " ") || strings.Contains(args[0], "/") {
			fmt.Println("Location name cannot contain spaces or slashes")
			return
		}

		existingLocations := viper.GetStringMapString("app.customlocations")
		existingLocations[args[0]] = args[1]
		viper.Set("app.customlocations", existingLocations)
		viper.WriteConfig()

		fmt.Println("Added custom location: " + args[0] + " -> " + args[1])
	},
}

var deleteLocationCmd = &cobra.Command{
	Use:     "delete <name>",
	Short:   "Delete a custom location",
	Example: "proto config locations delete steam",
	Aliases: []string{"del", "remove", "rm"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		existingLocations := viper.GetStringMapString("app.customlocations")
		if _, ok := existingLocations[args[0]]; ok {
			delete(existingLocations, args[0])
			viper.Set("app.customlocations", existingLocations)
			viper.WriteConfig()
			fmt.Println("Deleted custom location: " + args[0])
		} else {
			fmt.Println("That custom location does not exist")
		}
	},
}

var listLocationsCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all custom locations",
	Example: "proto config locations list",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		locations := viper.GetStringMapString("app.customlocations")
		if len(locations) == 0 {
			fmt.Println("No custom locations have been added.")
			return
		}

		for key, value := range locations {
			fmt.Println(key, "=", value)
		}
	},
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the configuration to default",
	Long:  "Reset the configuration to default, useful for when a major update occurs or you want to go back to defaults",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		os.Remove(viper.ConfigFileUsed())
		fmt.Println("Configuration has been reset to default.")
	},
}

var sourcesCmd = &cobra.Command{
	Use:   "sources <cmd>",
	Short: "Modify the sources list",
	Long:  `Sources are public GitHub repositories that Proto uses to find releases for you. Make sure you only add sources that you trust.`,
	Args:  cobra.ExactArgs(1),
}

var addSourceCmd = &cobra.Command{
	Use:     "add <owner/repo>",
	Short:   "Add a source to the list",
	Example: "add GloriousEggroll/proton-ge-custom",
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
		fmt.Println("Added source: " + args[0])
	},
}

var delSourceCmd = &cobra.Command{
	Use:     "del <owner/repo>",
	Short:   "Remove a source from the list",
	Example: "del GloriousEggroll/proton-ge-custom",
	Aliases: []string{"del", "remove", "rm"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var sources = viper.GetStringSlice("app.sources")

		for i, source := range sources {
			if source == args[0] {
				sources = append(sources[:i], sources[i+1:]...)
				break
			}
		}

		viper.Set("app.sources", sources)
		viper.WriteConfig()
		fmt.Println("Removed source: " + args[0])
	},
}

var listSourcesCmd = &cobra.Command{
	Use:   "list",
	Short: "List all sources",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		sources := viper.GetStringSlice("app.sources")
		if len(sources) == 0 {
			fmt.Println("No sources have been added.")
			return
		}

		fmt.Println("Currently configured sources:")
		for _, v := range sources {
			fmt.Println("- " + v)
		}
	},
}

func init() {
	RootCmd.AddCommand(configCmd)

	configCmd.AddCommand(showConfCmd)
	configCmd.AddCommand(configDirCmd)
	configCmd.AddCommand(tempCmd)
	configCmd.AddCommand(forceCmd)
	configCmd.AddCommand(verboseCmd)
	configCmd.AddCommand(sourcesCmd)
	configCmd.AddCommand(locationsCmd)
	configCmd.AddCommand(resetCmd)

	sourcesCmd.AddCommand(addSourceCmd)
	sourcesCmd.AddCommand(delSourceCmd)
	sourcesCmd.AddCommand(listSourcesCmd)

	locationsCmd.AddCommand(addLocationCmd)
	locationsCmd.AddCommand(deleteLocationCmd)
	locationsCmd.AddCommand(listLocationsCmd)
}
