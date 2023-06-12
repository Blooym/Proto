/*
Copyright © 2022 Blooym

GPLv3 License, see the LICENSE file for more information.
*/
package core

import (
	"context"
	"fmt"
	"os"
	"strings"

	github "github.com/google/go-github/v44/github"
	"github.com/spf13/viper"
)

/*
	FormatRepo takes a source index and returns the owner and repo.
	Arguments:
		entryIndex<int>: The index of the source to get the owner and repo from.
	Example:
		owner, repo := FormatRepo(0)
		fmt.Println(owner) // Blooym
		fmt.Println(repo) // proto
	Returns:
		string: The owner of the repo.
		string: The name of the repo.
*/
func FormatRepo(entryIndex int) (string, string) {

	sources := viper.GetStringSlice("app.sources")

	if len(sources) == 0 {
		fmt.Println("No sources have been configured. Please add a source with `proto config sources add <owner/repo>`.")
		os.Exit(1)
	}

	split := strings.Split(sources[entryIndex], "/")

	return split[0], split[1]
}

/*
	PromptSourceIndex asks the user which source they want to use if they do not manually specify one.
	Example:
		index := PromptSourceIndex()
		fmt.Println(index) // 0
	Returns:
		int: The index of the source the user selected.
*/
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

/*
	GetReleases returns all of the releases for the specified source index.
	Arguments:
		entryIndex<int>: The index of the source to get the owner and repo from.
	Example:
		releases, err := GetReleases(0)
	Returns:
		[]*github.RepositoryRelease: A list of all of the releases for the specified source index.
		error: Any errors that occur.
*/
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

/*
	GetReleaseData returns the release data for the specified source index and tag.
	Arguments:
		entryIndex<int>: The index of the source to get the owner and repo from.
		tag<string>: The tag of the release to get the data for.
	Example:
		release, err := GetReleaseData(0, "v1.0.0")
	Returns:
		*github.RepositoryRelease: The release data for the specified source index and tag.
		error: Any errors that occur.
*/
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

/*
	GetTotalAssetSize returns the total size of all of the assets in the specified release.
	Arguments:
		assets<[]*github.ReleaseAsset>: The assets to get the total size of.
	Example:
		size := GetTotalAssetSize(release.Assets)
		fmt.Println(size) // 123456789
	Returns:
		int64: The total size of all of the assets in the specified release.
*/
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

/*
	GetValidAssets returns a tar file and a sha512sum file from the specified release.
	Arguments:
		release<*github.RepositoryRelease>: The release to get the assets from.
	Example:
		assets := GetValidAssets(release.Assets)
		fmt.Println(assets) // [tar.xz, sha512sum]
	Returns:
		[]*github.ReleaseAsset: A list of the tar file and the sha512sum file.
*/
func GetValidAssets(release *github.RepositoryRelease) (*github.ReleaseAsset, *github.ReleaseAsset, error) {
	var runnerTar *github.ReleaseAsset
	var runnerSum *github.ReleaseAsset

	for _, asset := range release.Assets {

		Debug("GetValidAssets: Validating asset: " + asset.GetName())

		// Once we have both assets, we don't need to keep looking.
		if runnerTar != nil && runnerSum != nil {
			Debug("GetValidAssets: Found both assets, finishing search.")
			break
		}

		// Find the files needed for installing the runner.
		// Any tar file is supported, but it is recommended to use the .tar.xz format for better compression.
		if strings.HasSuffix(asset.GetName(), ".tar.gz") {
			Debug("GetValidAssets: Found a valid tar.gz asset.")
			runnerTar = asset
		} else if strings.HasSuffix(asset.GetName(), ".tar.xz") {
			Debug("GetValidAssets: Found a valid tar.xz asset.")
			runnerTar = asset
		} else if strings.HasSuffix(asset.GetName(), ".sha512sum") {
			Debug("GetValidAssets: Found a valid sha512sum asset.")
			runnerSum = asset
		}
	}

	// There was no tarball found for the release.
	if runnerTar == nil {
		return nil, nil, fmt.Errorf("unable to find a runner tarball")
	}

	// There was no valid checksum found for the release.
	if runnerSum == nil {
		return runnerTar, nil, nil
	}

	return runnerTar, runnerSum, nil
}
