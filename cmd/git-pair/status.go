package main

import (
	"fmt"
	"github.com/inverse/git-pair/internal/git"
)

func Status() {
	if !git.IsGitRepo() {
		fmt.Println("Not executed from a git repository")
		return
	}

	if !git.PairingModeEnabled() {
		fmt.Println("Pairing mode is not enabled")
		return
	}

	fmt.Println("Pairing mode is enabled")
}
