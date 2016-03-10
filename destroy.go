package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/teemow/loopback/image"
	"github.com/teemow/loopback/loop"
)

var (
	destroyCmd = &cobra.Command{
		Use:   "destroy",
		Short: "Destroy a loopback device",
		Long: `Destroy a loopback device.

Loopback is unmounted, deattached and images are removed in /var/lib/loopback.
`,
		Run: destroyRun,
	}

	destroyFlags struct {
		name      string
		imagePath string
	}
)

func init() {
	destroyCmd.Flags().StringVar(&destroyFlags.name, "name", "", "Name of the volume")
	destroyCmd.Flags().StringVar(&destroyFlags.imagePath, "image-path", "/var/lib/loopback", "Path for the loopback images")
}

func destroyRun(cmd *cobra.Command, args []string) {
	if destroyFlags.name == "" {
		fmt.Fprintln(os.Stderr, "Image name parameter missing.")
		os.Exit(1)
	}

	device, err := loop.Find(destroyFlags.name, destroyFlags.imagePath)
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
		fmt.Fprintf(os.Stderr, "Couldn't destroy loopback: %s\n", err)
		os.Exit(1)
	}

	err = image.Destroy(destroyFlags.name, destroyFlags.imagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't destroy image: %s\n", err)
		os.Exit(1)
	}
}
