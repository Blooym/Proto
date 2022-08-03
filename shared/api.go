/*
Copyright © 2022 BitsOfAByte

*/
package shared

import (
	"context"
	"fmt"
	"os"
	"strings"

	github "github.com/google/go-github/v44/github"
	"github.com/spf13/viper"
)

// Format the the given index source into a valid format.
func FormatRepo(entryIndex int) (string, string) {

	sources := viper.GetStringSlice("app.sources")

	if len(sources) == 0 {
		fmt.Println("No sources have been configured. Please add a source with `proto config sources add <owner/repo>`.")
		os.Exit(1)
	}

	split := strings.Split(sources[entryIndex], "/")

	return split[0], split[1]
}

// Get the source index the user is trying to install from.
func PromptSourceIndex() int {
	var source int
	sources := viper.GetStringSlice("app.sources")

	// if there is more than one source, ask the user which one they want to install from.
	if len(sources) > 1 {
		Debug("GetSourceIndex: Found " + fmt.Sprintf("%d", len(sources)) + " sources.")

		fmt.Println("\nMultiple sources found. Which one do you want to use?")
		for i, source := range sources {
			fmt.Printf("%d. %s\n", i+1, source)
		}
		fmt.Println("0. Cancel")
		fmt.Print("Choice: ")
		fmt.Scanf("%d", &source)

		// If the user cancels, exit.
		if source == 0 {
			os.Exit(0)
		}

		// If the user selects a source that doesn't exist, try again.
		if source < 1 || source > len(sources) {
			Debug("GetSourceIndex: User chose an invalid source.")
			return PromptSourceIndex()
		}

		// If the user selects a source that does exist, return the index minus one.
		fmt.Println("")
		Debug("GetSourceIndex: User chose source: " + sources[source-1])
		return source - 1
	}

	// If there is only one source, return the index.
	return 0
}

// Get all of the releases from the proton source.
func GetReleases(entryIndex int) ([]*github.RepositoryRelease, error) {
	owner, repo := FormatRepo(entryIndex)
	client := github.NewClient(nil)

	releases, _, err := client.Repositories.ListReleases(context.Background(), owner, repo, nil)

	Debug("GetReleases: Found " + fmt.Sprintf("%d", len(releases)) + " releases for " + owner + "/" + repo)

	if err != nil {
		return nil, err
	}

	return releases, nil
}

// Fetch all of the data for a specified tag from the proton source.
func GetReleaseData(entryIndex int, tag string) (*github.RepositoryRelease, error) {
	owner, repo := FormatRepo(entryIndex)
	client := github.NewClient(nil)

	release, _, err := client.Repositories.GetReleaseByTag(context.Background(), owner, repo, tag)

	Debug("GetReleaseData: Looking for: " + owner + "/" + repo + "/" + tag)

	if err != nil {
		return nil, err
	}

	Debug("GetReleaseData: Found release " + release.GetTagName())

	return release, nil
}

// Total Asset size in bytes.
func GetTotalAssetSize(assets []*github.ReleaseAsset) int64 {
	var size int

	// Loop through all of the assets and add their sizes together.
	for _, asset := range assets {
		if strings.HasSuffix(asset.GetName(), ".tar.gz") {
			size += int(asset.GetSize())
		}

		if strings.HasSuffix(asset.GetName(), ".tar.xz") {
			size += int(asset.GetSize())
		}

		if strings.HasSuffix(asset.GetName(), ".sha512sum") {
			size += asset.GetSize()
		}
	}

	return int64(size)
}

// Sorts through a release to find valid assets for downloading a release.
func GetValidAssets(release *github.RepositoryRelease) (*github.ReleaseAsset, *github.ReleaseAsset, error) {
	var protonTar *github.ReleaseAsset
	var protonSum *github.ReleaseAsset

	for _, asset := range release.Assets {

		Debug("GetValidAssets: Validating asset: " + asset.GetName())

		// Once we have both assets, we don't need to keep looking.
		if protonTar != nil && protonSum != nil {
			Debug("GetValidAssets: Found both assets, finishing search.")
			break
		}

		// Find the files needed for installing the proton.
		// Any tar file is supported, but it is recommended to use the .tar.xz format for better compression.
		if strings.HasSuffix(asset.GetName(), ".tar.gz") {
			Debug("GetValidAssets: Found a valid tar.gz asset.")
			protonTar = asset
		} else if strings.HasSuffix(asset.GetName(), ".tar.xz") {
			Debug("GetValidAssets: Found a valid tar.xz asset.")
			protonTar = asset
		} else if strings.HasSuffix(asset.GetName(), ".sha512sum") {
			Debug("GetValidAssets: Found a valid sha512sum asset.")
			protonSum = asset
		}
	}

	// There was no tarball found for the release.
	if protonTar == nil {
		return nil, nil, fmt.Errorf("unable to find a proton tarball")
	}

	// There was no valid checksum found for the release.
	if protonSum == nil {
		return protonTar, nil, nil
	}

	return protonTar, protonSum, nil
}
