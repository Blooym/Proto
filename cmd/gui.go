/*
Copyright © 2022 BitsOfAByte

*/
package cmd

import (
	"BitsOfAByte/proto/shared"
	"fmt"

	"github.com/spf13/cobra"
)

// guiCmd represents the gui command
var guiCmd = &cobra.Command{
	Use:   "gui",
	Short: "Opens the Proto GUI (Coming Soon)",
	Run: func(cmd *cobra.Command, args []string) {
		lock := shared.HandleLock()
		defer lock.Unlock()

		fmt.Println("The GUI has not been implemented yet! For now you'll have to use the command line.")
		fmt.Println("For help with the command line, use the 'proto help' command.")
	},
}

func init() {
	rootCmd.AddCommand(guiCmd)
}
