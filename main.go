/*
Copyright © 2022 BitsOfAByte

*/
package main

import (
	"fmt"

	"github.com/BitsOfAByte/proto/cmd"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("FATAL: %v", r)
		}
	}()

	cmd.Execute()
}
