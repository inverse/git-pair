package main

import (
	"fmt"

	"github.com/inverse/git-pair/internal/git"
)

func End() {
	if !git.IsGitRepo() {
		fmt.Println("Not executed from a git repository")
		return
	}

	err := git.DisablePairingMode()

	if err != nil {
		fmt.Println(err)
		return
	}
}
