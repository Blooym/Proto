/*
Copyright © 2022 BitsOfAByte

*/
package cmd

import (
	"BitsOfAByte/proto/shared"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info <tag>",
	Short: "Shows information about the given release.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		// If there are multiple sources, ask the user which one to use or use the flag.
		var source int
		sourceFlag, _ := cmd.Flags().GetInt("source")
		if sourceFlag > 0 {
			sourceLen := len(viper.GetStringSlice("app.sources"))

			if sourceFlag-1 >= sourceLen {
				fmt.Println("There is no source at index", sourceFlag, "you only have", sourceLen, "sources.")
				os.Exit(1)
			}

			source = sourceFlag - 1
		} else {
			source = shared.PromptSourceIndex()
		}

		// Fetch the release data.
		data, err := shared.GetReleaseData(source, args[0])

		if err != nil {
			fmt.Println("That release does not exist on the given source.")
			os.Exit(1)
		}

		// Print the release data.
		fmt.Println("Release:", data.GetTagName())
		fmt.Println("Published:", data.GetPublishedAt().Format("2006-01-02 15:04:05"))
		fmt.Println("Description:", data.GetBody())
		fmt.Println("Install Command: proto install", data.GetTagName(), "-s", source+1)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	infoCmd.Flags().IntP("source", "s", 0, "The index of the source to use.")
}
