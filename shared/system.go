/*
Copyright © 2022 BitsOfAByte

*/
package shared

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofrs/flock"
)

// Restricts running the binary as a root user. (Identified by UID == 0)
func IsRoot() bool {
	return os.Geteuid() == 0
}

// Detects if the binary is a preview version. (Identified by Version string)
func IsPreview() bool {
	return Version == "0.0.1-next" && os.Getenv("PROTO_CONSENT") != "true"
}

// Handles the locking and unlocking of the program lock file.
func HandleLock() *flock.Flock {
	// Get the file lock.
	fileLock := flock.New("/tmp/proto.lock")
	locked, err := fileLock.TryLock()
	Check(err)

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

// Prompt a user to confirm the given input with a yes/no prompt.
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
