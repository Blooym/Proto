/*
Copyright © 2022 BitsOfAByte

*/
package cmd

import (
	"BitsOfAByte/proto/shared"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// appUpdateCmd represents the update command
var appUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update to the latest version of Proto",
	Run: func(cmd *cobra.Command, args []string) {
		forceFlag := cmd.Flag("force").Value.String()
		if forceFlag == "true" {
			lock := shared.HandleLock()
			defer lock.Unlock()

			shared.AppUpdate(shared.Version)
		} else {
			fmt.Println("WARNING! You should not use the app-update command unless you have a manual installation of Proto.")
			fmt.Println("If you are trying to update the app and have installed it with a package manager, use that instead.")
			fmt.Println("If you are ABSOLUTELY sure you want to update Proto, use the --force flag.")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(appUpdateCmd)

	appUpdateCmd.Flags().BoolP("force", "f", false, "Force the updater to run")
}
