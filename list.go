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
)

func listRun(cmd *cobra.Command, args []string) {
	list, err := loop.List()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't list devices: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(list)
	os.Exit(0)
}
