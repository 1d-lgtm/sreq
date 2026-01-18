package main

import (
	"os"

	"github.com/Priyans-hu/sreq/cmd/sreq/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
