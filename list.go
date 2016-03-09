package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/teemow/loopback/loop"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List loopback devices",
		Long: `List loopback devices.

List of loopback devices on a system
`,
		Run: listRun,
	}

	listFlags struct {
		noHeadings bool
	}
)

func init() {
	listCmd.Flags().BoolVar(&listFlags.noHeadings, "no-headings", false, "Don't print headings on top of the list")
}

func listRun(cmd *cobra.Command, args []string) {
	var err error
	var list string

	if listFlags.noHeadings {
		list, err = loop.ListWithoutHeadings()
	} else {
		list, err = loop.List()
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't list devices: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(list)
	os.Exit(0)
}
