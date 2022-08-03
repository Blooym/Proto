/*
Copyright © 2022 BitsOfAByte

*/
package cmd

import (
	"BitsOfAByte/proto/shared"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:        "uninstall <version>",
	Short:      "Uninstall a version of Proton from your system.",
	Aliases:    []string{"rm", "remove"},
	SuggestFor: []string{"delete"},
	Example:    "proto uninstall GE-Proton7-18",
	Args:       cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		// Prevent the program from having another long-running process
		lock := shared.HandleLock()
		defer lock.Unlock()

		installDir := shared.UsePath(viper.GetString("storage.install"), true) + args[0]

		if _, err := os.Stat(installDir); os.IsNotExist(err) {
			fmt.Println("The specified version of Proton was not found at " + filepath.Dir(installDir))
			os.Exit(1)
		}

		// Prompt the user to confirm unless -y flag is set.
		yesFlag := rootCmd.Flag("yes").Value.String()
		if yesFlag != "true" {
			// Prompt the user to confirm the uninstall.
			resp := shared.Prompt("Are you sure you want to uninstall Proton "+args[0]+"? (y/N) ", false)

			if !resp {
				os.Exit(0)
			}
		}

		// Remove the directory for the specified version.
		err := os.RemoveAll(installDir)
		shared.Check(err)

		fmt.Printf("Successfully uninstalled %s from %s\n", args[0], filepath.Dir(installDir))
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
