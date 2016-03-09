package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/teemow/loopback/loop"
)

var (
	detachCmd = &cobra.Command{
		Use:   "detach",
		Short: "Detach a loopback device",
		Long: `Detach a loopback device.

Loopback is deattached.
`,
		Run: detachRun,
	}

	detachFlags struct {
		name      string
		imagePath string
	}
)

func init() {
	detachCmd.Flags().StringVar(&detachFlags.name, "name", "", "Name of the volume")
	detachCmd.Flags().StringVar(&detachFlags.imagePath, "image-path", "/var/lib/loopback", "Path for the loopback images")
}

func detachRun(cmd *cobra.Command, args []string) {
	if detachFlags.name == "" {
		fmt.Fprintln(os.Stderr, "Image name parameter missing.")
		os.Exit(1)
	}

	device, err := loop.Find(detachFlags.name, detachFlags.imagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't find loopback: %s\n", err)
		os.Exit(1)
	}

	err = loop.Destroy(device)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't destroy loopback: %s\n", err)
		os.Exit(1)
	}
}
