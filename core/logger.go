/*
Copyright © 2022 BitsOfAByte

GPLv3 License, see the LICENSE file for more information.
*/
package core

import (
	"fmt"

	"github.com/spf13/viper"
)

/*
	Debug is a function that prints a debug message to the console if the verbose flag is set.
	Arguments:
		msg<string>: The message to print
	Example:
		Debug("This is a debug message.")
*/
func Debug(msg string) {
	if viper.GetBool("cli.verbose") {
		fmt.Printf("[DEBUG] %s\n", msg)
	}
}

/*
	CheckError is a function that checks if an error is nil. If it is not, it will panic.
	Arguments:
		err<error>: The error to check
	Example:
		CheckError(err)
*/
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
