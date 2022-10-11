/*
Copyright © 2022 BitsOfAByte

GPLv3 License, see the LICENSE file for more information.
*/
package cmd

import (
	"BitsOfAByte/proto/shared"
	"fmt"
	"os"

	"github.com/google/go-github/v44/github"
	cobra "github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var installCmd = &cobra.Command{
	Use:   "install [tag]",
	Short: "Download and install runner to your system.",
	PostRun: func(cmd *cobra.Command, args []string) {
		shared.ClearTemp()
	},
	Run: func(cmd *cobra.Command, args []string) {

		// Prevent the program from having another long-running process
		lock := shared.HandleLock()
		defer lock.Unlock()

		// Make sure an install directory is specified.
		installDir := cmd.Flags().Lookup("dir").Value.String()
		if installDir == "" {
			fmt.Println("No operating directory specified, please use the --dir flag to specify either a full path or a custom keyword path (run 'proto config locations -h' for more info).")
			os.Exit(1)
		}
		installDir = shared.UsePath(shared.GetCustomLocation(installDir), true)

		/**
		----------------------
		|     Fetch Logic    |
		----------------------
		**/

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

		// Find the version to install, if none is specified, use the latest.
		var tagData *github.RepositoryRelease
		switch len(args) {
		case 0: // Install latest tag.
			data, err := shared.GetReleases(source)
			shared.CheckError(err)
			tagData = data[0]
		default: // Install a specific tag.
			data, err := shared.GetReleaseData(source, args[0])
			shared.CheckError(err)
			tagData = data
		}

		yesFlag := rootCmd.Flag("yes").Value.String()
		s, m := shared.HumanReadableBytes(shared.GetTotalAssetSize(tagData.Assets))

		/**
		----------------------
		|    Overlap Logic   |
		----------------------
		**/

		// Check if the directory exists already, meaning we're trying to install a version that's already installed.
		if folderInfo, err := os.Stat(installDir + tagData.GetTagName()); err == nil && folderInfo.IsDir() {
			// Prompt the user for to overwrite the existing version, skipped if -y flag is set.
			if yesFlag != "true" {
				resp := shared.Prompt(fmt.Sprintf("Looks like %s is already installed, overwrite? [Est. %v%s] (y/N) ", tagData.GetTagName(), s, m), false)

				if !resp {
					os.Exit(0)
				}
			}

			// Remove the existing directory.
			if err := os.RemoveAll(installDir + tagData.GetTagName()); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println("Removed old installation: " + tagData.GetTagName())
		} else if yesFlag != "true" {
			// Prompt the user to confirm the install, skipped if -y flag is set.
			resp := shared.Prompt(fmt.Sprintf("Are you sure you want to install %s? [Est. %v%s] (y/N) ", tagData.GetTagName(), s, m), false)

			if !resp {
				os.Exit(0)
			}
		}

		/**
		----------------------
		|   Download Logic   |
		----------------------
		**/

		// Fetch valid assets from the release.
		tar, sum, err := shared.GetValidAssets(tagData)
		shared.CheckError(err)

		// Handle the lack of a checksum depending on the user's preference.
		if sum == nil {
			forced := viper.GetBool("app.forcechecksum")

			if !forced {
				fmt.Println("No checksum file found, aborting install. (Use --force to ignore this error this time)")
				os.Exit(1)
			}

			fmt.Println("Warning: No checksum file found, continuing without verification (forced).")
		}

		// Download the assets to the temp directory.
		tmp := shared.UsePath(viper.GetString("storage.tmp"), true)

		// Download the tarball.
		_, err = shared.DownloadFile(tmp+tar.GetName(), tar.GetBrowserDownloadURL())
		shared.CheckError(err)

		/**
		----------------------
		|   Checksum Logic   |
		----------------------
		**/

		// If it exists, download the checksum file and verify it against the downloaded tarball.
		if sum != nil {
			_, err = shared.DownloadFile(tmp+sum.GetName(), sum.GetBrowserDownloadURL())
			shared.CheckError(err)

			match, err := shared.MatchChecksum(tmp+tar.GetName(), tmp+sum.GetName())
			forceSum := viper.GetBool("app.forcechecksum")

			shared.CheckError(err)

			// If the checksums don't match and force sum is enabled, abort.
			if !match && viper.GetBool("app.forcechecksum") {
				fmt.Println("Checksums do not match, aborting install.")
				os.Exit(1)
			}

			// If the checksums don't match and force sum is disabled, prompt the user to continue unless -y flag is set.
			if !match && !forceSum && yesFlag != "true" {
				resp := shared.Prompt(fmt.Sprintf("Checksums do not match, continue? [Est. %v%s] (y/N) ", s, m), false)

				if !resp {
					os.Exit(0)
				}

			} else if !match && !forceSum && yesFlag == "true" {
				// -y flag is set, warn the user that the checksums don't match.
				fmt.Println("Warning! Checksums do not match, continuing without verification due to -y flag.")
			}

			// Everything checks out, continue with the install.
			if match {
				fmt.Println("Checksums verified successfully.")
			}
		}

		/**
		----------------------
		|   Install Logic    |
		----------------------
		**/

		fmt.Println("Extracting files...")

		err = shared.ExtractTar(tmp+tar.GetName(), installDir)
		shared.CheckError(err)

		/**
		----------------------
		| Post-Install Logic |
		----------------------
		**/

		fmt.Printf("%s has been successfully installed!\nLocation: %s\n", tagData.GetTagName(), installDir)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Register the command flags.
	installCmd.Flags().StringP("dir", "d", "", "The directory to operate in.")
	installCmd.Flags().BoolP("force", "f", false, "Force installation (ignoring missing or failed checksums)")
	installCmd.Flags().IntP("source", "s", 0, "Specify the source to install from.")

	// Bind the flags to the viper config.
	viper.BindPFlag("app.forcechecksum", installCmd.Flags().Lookup("force"))
}
