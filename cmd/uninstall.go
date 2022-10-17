/*
Copyright © 2022 BitsOfAByte

GPLv3 License, see the LICENSE file for more information.
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BitsOfAByte/proto/core"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:        "uninstall <version>",
	Short:      "Uninstall a runner from your system.",
	Aliases:    []string{"rm", "remove"},
	SuggestFor: []string{"delete"},
	Example:    "proto uninstall GE-Proton7-18 --dir ~/.steam/root/compatibilitytools.d/",
	Args:       cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		// Prevent the program from having another long-running process
		lock := core.HandleLock()
		defer lock.Unlock()

		getDir := cmd.Flags().Lookup("dir").Value.String()
		if getDir == "" {
			fmt.Println("No operating directory specified, please use the --dir flag to specify either a full path or a custom keyword path (run 'proto config locations -h' for more info).")
			os.Exit(1)
		}
		getDir = core.UsePath(core.GetCustomLocation(getDir), true) + args[0]

		if _, err := os.Stat(getDir); os.IsNotExist(err) {
			fmt.Println("The specified runner was not found at " + filepath.Dir(getDir))
			os.Exit(1)
		}

		// Prompt the user to confirm unless -y flag is set.
		yesFlag := RootCmd.Flag("yes").Value.String()
		if yesFlag != "true" {
			// Prompt the user to confirm the uninstall.
			resp := core.Prompt("Are you sure you want to uninstall the runner "+args[0]+"? (y/N) ", false)

			if !resp {
				os.Exit(0)
			}
		}

		// Remove the directory for the specified version.
		err := os.RemoveAll(getDir)
		core.CheckError(err)

		fmt.Printf("Successfully uninstalled %s from %s\n", args[0], filepath.Dir(getDir))
	},
}

func init() {
	RootCmd.AddCommand(uninstallCmd)
}
