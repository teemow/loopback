package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	globalFlags struct {
		debug   bool
		verbose bool
	}

	LoopbackCmd = &cobra.Command{
		Use:   "loopback",
		Short: "Manage loopback devices",
		Long:  "Manage loopback devices",
		Run:   LoopbackRun,
	}

	projectVersion string
)

func init() {
	LoopbackCmd.PersistentFlags().BoolVarP(&globalFlags.debug, "debug", "d", false, "Print debug output")
	LoopbackCmd.PersistentFlags().BoolVarP(&globalFlags.verbose, "verbose", "v", false, "Print verbose output")
}

func assert(err error) {
	if err != nil {
		if globalFlags.debug {
			fmt.Printf("%#v\n", err)
			os.Exit(1)
		} else {
			log.Fatal(err)
		}
	}
}

func confirm(question string) error {
	for {
		fmt.Printf("%s ", question)
		bio := bufio.NewReader(os.Stdin)
		line, _, err := bio.ReadLine()
		if err != nil {
			return err
		}

		if string(line) == "yes" {
			return nil
		}
		fmt.Println("Please enter 'yes' to confirm.")
	}
}

func LoopbackRun(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func main() {
	LoopbackCmd.AddCommand(versionCmd)
	LoopbackCmd.AddCommand(createCmd)
	LoopbackCmd.AddCommand(listCmd)
	LoopbackCmd.AddCommand(listImagesCmd)
	LoopbackCmd.AddCommand(destroyCmd)

	LoopbackCmd.Execute()
}
