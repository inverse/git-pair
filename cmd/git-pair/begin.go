package main

import (
	"fmt"

	"github.com/inverse/git-pair/internal/git"
)

func Begin() {
	if !git.IsGitRepo() {
		fmt.Println("Not executed from a git repository")
		return
	}

	git.GetRepoContributors()

	contributors := []string{"Malachi Soord <inverse.chi@gmail.com"}

	git.EnablePairingMode(contributors)
}
