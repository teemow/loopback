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
		mountPath string
	}
)

func init() {
	createCmd.Flags().StringVar(&createFlags.name, "name", "", "Name of the volume")
	createCmd.Flags().StringVar(&createFlags.imagePath, "image-path", "/var/lib/loopback", "Path for the loopback images")
	createCmd.Flags().IntVar(&createFlags.size, "size", 1, "Size of the volume (in gigabytes)")
	createCmd.Flags().StringVar(&createFlags.fs, "fs", "btrfs", "Filesystem")
	createCmd.Flags().StringVar(&createFlags.mountPath, "mount-path", "", "Path to mount loopback device into")
}

func createRun(cmd *cobra.Command, args []string) {
	var err error

	if createFlags.mountPath != "" {
		_, err := os.Stat(createFlags.mountPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Mount path does not exist.")
			os.Exit(1)
		}
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

	if createFlags.mountPath != "" {
		loop.Mount(device, createFlags.mountPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't mount loopback: %s\n", err)
			os.Exit(1)
		}
	}

	os.Exit(0)
}
