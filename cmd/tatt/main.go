package main

import (
	"fmt"
	"os"

	"github.com/michenriksen/tatt/internal/cli"
)

func main() {
	if err := cli.App().Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
