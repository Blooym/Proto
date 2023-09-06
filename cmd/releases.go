package cmd

import (
	"fmt"
	"os"

	"github.com/Blooym/proto/core"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var releasesCmd = &cobra.Command{
	Use:     "releases",
	Short:   "Show all available releases from the runner source.",
	Example: `proto releases --limit 5`,
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
			source = core.PromptSourceIndex()
		}

		// Get the releases from the backend.
		releases, err := core.GetReleases(source)
		core.CheckError(err)

		// Create a table to display the releases.
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Tag", "Released On", "Info Command"})
		limit, _ := cmd.Flags().GetInt("limit")

		// Loop through the releases and add them to the table up to the limit.
		for _, release := range releases {

			if limit > 0 {
				limit--
			} else {
				break
			}

			table.Append([]string{
				release.GetTagName(),
				release.GetPublishedAt().Format("2006-01-02"),
				fmt.Sprintf("proto info %s -s %d", release.GetTagName(), source+1),
			})
		}

		// Display the table.
		table.Render()
	},
}

func init() {
	RootCmd.AddCommand(releasesCmd)

	// Register command flags
	releasesCmd.Flags().IntP("limit", "l", 5, "Limit the number of releases to show.")
	releasesCmd.Flags().IntP("source", "s", 0, "The source to use.")
}
