package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/teemow/loopback/loop"
)

var (
	mountCmd = &cobra.Command{
		Use:   "mount",
		Short: "Mount a loopback device",
		Long: `Mount a loopback device.

Attach an image as a loopback and then mount it.
`,
		Run: mountRun,
	}

	mountFlags struct {
		name      string
		imagePath string
		mountPath string
	}
)

func init() {
	mountCmd.Flags().StringVar(&mountFlags.name, "name", "", "Name of the volume")
	mountCmd.Flags().StringVar(&mountFlags.imagePath, "image-path", "/var/lib/loopback", "Path for the loopback images")
	mountCmd.Flags().StringVar(&createFlags.mountPath, "mount-path", "", "Path to mount loopback device into")
}

func mountRun(cmd *cobra.Command, args []string) {
	if mountFlags.mountPath == "" {
		fmt.Fprintln(os.Stderr, "Mount path parameter missing.")
		os.Exit(1)
	}

	_, err := os.Stat(mountFlags.mountPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Mount path does not exist.")
		os.Exit(1)
	}

	var device string
	device, err = loop.Create(mountFlags.name, mountFlags.imagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't attach loopback: %s\n", err)
		os.Exit(1)
	}

	loop.Mount(device, mountFlags.mountPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't mount loopback: %s\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
