package main

import (
	"fmt"
	"github.com/inverse/git-pair/internal/diagnostics"
)

func main() {
	printVersion()
}

func printVersion() {
	fmt.Printf(
		"\nversion: %s\ncommit: %s\nbuild time: %s\n",
		diagnostics.Version, diagnostics.Commit, diagnostics.BuildDate,
	)
}
