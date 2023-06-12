/*
Copyright © 2022 Blooym

GPLv3 License, see the LICENSE file for more information.
*/
package core

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofrs/flock"
)

/*
	HandleLock is a function that handles the file lock for Proto preventing multiple instances of the app from running at once.
	Example:
		lock := HandleLock()
		defer lock.Unlock()
	Returns:
		flock.Flock: The file lock
*/
func HandleLock() *flock.Flock {
	// Create a cache directory if it doesn't exist
	cacheDir, _ := os.UserCacheDir()
	os.MkdirAll(cacheDir+"/proto", 0755)

	// Create a lock file
	fileLock := flock.New(cacheDir + "/proto/lockfile")
	locked, err := fileLock.TryLock()
	CheckError(err)

	// The lock has been acquired, safe to proceed.
	if locked {
		Debug("Lock: Successfully acquired lock")
		return fileLock
	}

	// The lock is held by another process, exit.
	Debug("Lock: Failed to acquire lock, is the process already running?")
	fmt.Println("Another instance of Proto is already running, please close it and try again.")
	os.Exit(1)

	// Return nil to satisfy the compiler.
	return nil
}

/*
	Prompt is a function that prompts the user for a yes or no answer with a given message.
	Arguments:
		message<string> The message to display to the user.
		defaultValue<bool> The default value to return if the user hits enter
	Returns:
		bool: The user's answer
*/
func Prompt(message string, defaultValue bool) bool {
	var response string

	Debug("Prompt: Asking for user input")

	fmt.Print(message)
	fmt.Scanln(&response)

	switch strings.ToLower(response) {
	case "y", "yes":
		return true
	case "n", "no":
		return false
	default:
		return defaultValue
	}
}
