package main

import (
	"os"

	"github.com/Swing-Technologies/keplars-email-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
