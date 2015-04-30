package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/teemow/loopback/image"
	"github.com/teemow/loopback/loop"
)

var (
	createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a loopback device",
		Long: `Create a loopback device.

Images are created in /var/lib/loopback. Then attached to a loopback, formatted and then mounted.
`,
		Run: createRun,
	}

	createFlags struct {
		name      string
		imagePath string
		size      int
		fs        string
	}
)

func init() {
	createCmd.Flags().StringVar(&createFlags.name, "name", "", "Name of the volume")
	createCmd.Flags().StringVar(&createFlags.imagePath, "image-path", "/var/lib/loopback", "Path for the loopback images")
	createCmd.Flags().IntVar(&createFlags.size, "size", 1, "Size of the volume (in gigabytes)")
	createCmd.Flags().StringVar(&createFlags.fs, "fs", "btrfs", "Filesystem")
}

func createRun(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "Mount path parameter missing.")
		os.Exit(1)
	}

	mountPath := args[0]

	_, err := os.Stat(mountPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Mount path does not exist.")
		os.Exit(1)
	}

	err = image.Create(createFlags.name, createFlags.imagePath, createFlags.size)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't create image: %s\n", err)
		os.Exit(1)
	}

	var device string
	device, err = loop.Create(createFlags.name, createFlags.imagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't create loopback: %s\n", err)
		os.Exit(1)
	}

	loop.Format(device, createFlags.fs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't format loopback: %s\n", err)
		os.Exit(1)
	}

	loop.Mount(device, mountPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't mount loopback: %s\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
