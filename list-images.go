package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/teemow/loopback/image"
)

var (
	listImagesCmd = &cobra.Command{
		Use:   "list-images",
		Short: "List loopback images",
		Long: `List loopback images.

List images for loopback devices
`,
		Run: listImagesRun,
	}

	listImagesFlags struct {
		imagePath string
	}
)

func init() {
	listImagesCmd.Flags().StringVar(&listImagesFlags.imagePath, "image-path", "/var/lib/loopback", "Path for the loopback images")
}

func listImagesRun(cmd *cobra.Command, args []string) {
	list, err := image.List(listImagesFlags.imagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't list images: %s\n", err)
		os.Exit(1)
	}
	for _, image := range list {
		fmt.Println(image)
	}

	os.Exit(0)
}
