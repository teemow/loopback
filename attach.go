package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/teemow/loopback/image"
	"github.com/teemow/loopback/loop"
)

var (
	attachCmd = &cobra.Command{
		Use:   "attach",
		Short: "Attach a loopback device",
		Long: `Attach a loopback device.

Attach an image as a loopback.
`,
		Run: attachRun,
	}

	attachFlags struct {
		name               string
		imagePath          string
		size               int
		fs                 string
		createMissingImage bool
	}
)

func init() {
	attachCmd.Flags().StringVar(&attachFlags.name, "name", "", "Name of the volume")
	attachCmd.Flags().StringVar(&attachFlags.imagePath, "image-path", "/var/lib/loopback", "Path for the loopback images")
	attachCmd.Flags().IntVar(&attachFlags.size, "size", 1, "Size of the volume (in gigabytes)")
	attachCmd.Flags().StringVar(&attachFlags.fs, "fs", "btrfs", "Filesystem")
	attachCmd.Flags().BoolVar(&attachFlags.createMissingImage, "create-missing-image", false, "Creating image if it does not exist")
}

func attachRun(cmd *cobra.Command, args []string) {
	var err error

	if attachFlags.name == "" {
		fmt.Fprintln(os.Stderr, "Image name parameter missing.")
		os.Exit(1)
	}

	createAndFormatImage := false
	if !image.Exists(attachFlags.name, attachFlags.imagePath) || !attachFlags.createMissingImage {
		createAndFormatImage = true
	}

	if createAndFormatImage {
		err = image.Create(attachFlags.name, attachFlags.imagePath, attachFlags.size)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't create image: %s\n", err)
			os.Exit(1)
		}
	}

	var device string
	device, err = loop.Attach(attachFlags.name, attachFlags.imagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't attach loopback: %s\n", err)
		os.Exit(1)
	}

	if createAndFormatImage {
		loop.Format(device, attachFlags.fs)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't format loopback: %s\n", err)
			os.Exit(1)
		}
	}

	os.Exit(0)
}
