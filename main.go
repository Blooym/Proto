/*
Copyright © 2022 BitsOfAByte

*/
package main

import (
	"BitsOfAByte/proto/cmd"
	"BitsOfAByte/proto/shared"
	"fmt"
	"os"
)

func main() {

	if shared.IsRoot() {
		fmt.Println("Proto cannot not be run as root, please try again as a regular user without sudo.")
		os.Exit(1)
	}

	if shared.IsPreview() {
		fmt.Println("Detected a preview version of Proto! If you are a user, please download a stable version from GitHub directly.")
		fmt.Println("If you are certain you want to run this preview build, set environment variable 'PROTO_CONSENT' to 'true'")
		os.Exit(1)
	}

	cmd.Execute()
}
