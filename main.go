/*
Copyright © 2022 Blooym

*/
package main

import (
	"fmt"

	"github.com/Blooym/proto/cmd"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("FATAL: %v", r)
		}
	}()

	cmd.Execute()
}
