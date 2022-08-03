/*
Copyright © 2022 BitsOfAByte

*/
package shared

import (
	"fmt"

	"github.com/spf13/viper"
)

// Log a message if the logger is enabled.
func Debug(msg string) {
	if viper.GetBool("cli.verbose") {
		fmt.Printf("[DEBUG] %s\n", msg)
	}
}

// Checks the given error to see if its nil, if its not then panics.
func Check(err error) {
	if err != nil {
		panic(err)
	}
}
