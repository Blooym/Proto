/*
Copyright © 2022 Blooym

GPLv3 License, see the LICENSE file for more information.
*/
package core

import (
	"fmt"
	"os"
	"runtime"

	"github.com/creativeprojects/go-selfupdate"
)

/*
	AppUpdate is a function that checks for updates and updates the application if an update is available.
	Typically used for non package-manager installations
	Arguments:
		version<string> The current version of the application
	Example:
		AppUpdate("v1.2.3")
*/
func AppUpdate(version string) {
	updater, _ := selfupdate.NewUpdater(selfupdate.Config{Validator: &selfupdate.ChecksumValidator{UniqueFilename: "checksums.txt"}})
	latest, found, err := updater.DetectLatest("Blooym/proto")

	// Unknown error occurred, abort update process.
	CheckError(err)

	// Specified OS or Architechture is not supported.
	if !found {
		fmt.Printf("version %s is not supported on %s/%s", version, runtime.GOOS, runtime.GOARCH)
		return
	}

	// No update is available for the current version.
	if latest.LessOrEqual(version) {
		fmt.Printf("no update available for version %s\n", version)
		return
	}

	// Find the current executable's path.
	exe, err := os.Executable()

	//  Could not find the executable's path, abort update process.
	if err != nil {
		fmt.Printf("error occurred while finding executable's path: %v", err)
		return
	}

	// Perform the update.
	if err := selfupdate.UpdateTo(latest.AssetURL, latest.AssetName, exe); err != nil {
		fmt.Printf("error occurred while updating: %v", err)
		return
	}

	fmt.Printf("Successfully updated to version %s (OS: %s, Arch: %s) from %s\n", latest.Version(), latest.OS, latest.Arch, latest.PublishedAt)
}
