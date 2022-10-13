/*
Copyright © 2022 BitsOfAByte

*/
package main

import (
	"BitsOfAByte/proto/cmd"
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("FATAL: %v", r)
		}
	}()

	cmd.Execute()
}
