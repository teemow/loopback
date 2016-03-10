package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/teemow/loopback/loop"
)

var (
	umountCmd = &cobra.Command{
		Use:   "umount",
		Short: "Unmount a loopback device",
		Long: `Unmount a loopback device.

Unmount a loopback and then detach the device.
`,
		Run: umountRun,
	}

	umountFlags struct {
		name      string
		imagePath string
	}
)

func init() {
	umountCmd.Flags().StringVar(&umountFlags.name, "name", "", "Name of the volume")
	umountCmd.Flags().StringVar(&umountFlags.imagePath, "image-path", "/var/lib/loopback", "Path for the loopback images")
}

func umountRun(cmd *cobra.Command, args []string) {
	device, err := loop.Find(umountFlags.name, umountFlags.imagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't find loopback: %s\n", err)
		os.Exit(1)
	}

	err = loop.Unmount(device)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't unmount loopback: %s\n", err)
		os.Exit(1)
	}

	err = loop.Detach(device)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't detach loopback: %s\n", err)
		os.Exit(1)
	}
}
