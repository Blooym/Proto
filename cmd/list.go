/*
Copyright © 2022 BitsOfAByte

GPLv3 License, see the LICENSE file for more information.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BitsOfAByte/proto/core"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "Shows a list of installed runner versions.",
	Aliases: []string{"ls"},
	Example: "proto list --dir ~/.steam/root/compatibilitytools.d",
	Run: func(cmd *cobra.Command, args []string) {

		// Read the install directory
		getDir := cmd.Flag("dir").Value.String()
		if getDir == "" {
			fmt.Println("No operating directory specified, please use the --dir flag to specify either a full path or a custom keyword path (run 'proto config locations -h' for more info).")
			os.Exit(1)
		}
		getDir = core.UsePath(core.GetCustomLocation(getDir), true)

		dir, err := ioutil.ReadDir(getDir)
		if err != nil {
			// The directory doesnt exist, meaning there are no installed versions.
			if os.IsNotExist(err) {
				fmt.Println("No installed runners found at " + getDir)
				os.Exit(0)
			}

			// Something else went wrong, eg. permissions.
			fmt.Println(err)
			os.Exit(1)
		}

		// Get all of the installed versions and their sizes and create a table to display them.
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Version", "Size", "Installed", "Remove Command"})
		var totalSize int64
		for _, d := range dir {
			size, err := core.GetDirSize(getDir + d.Name())

			// Something went wrong getting the size of the directory.
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Get the size of the directory, and add to the total size, then append to the table.
			hSize, hUnit := core.HumanReadableBytes(size)
			totalSize += size
			table.Append([]string{d.Name(), fmt.Sprintf("%v%s", hSize, hUnit), d.ModTime().Format("2006-01-02"), fmt.Sprintf("proto uninstall %s --dir %s", d.Name(), getDir)})
		}

		// No installed versions found in the install directory.
		if table.NumLines() == 0 {
			fmt.Println("No installed runners found at " + getDir)
			os.Exit(0)
		}

		// Format the total size and render the table.
		tSize, tUnit := core.HumanReadableBytes(totalSize)
		table.SetFooter([]string{"Total", fmt.Sprintf("%v%s", tSize, tUnit), " ", " "})
		table.Render()
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
