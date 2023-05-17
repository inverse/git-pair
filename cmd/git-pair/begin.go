package main

import (
	"fmt"

	"github.com/inverse/git-pair/internal/contributors"
	"github.com/inverse/git-pair/internal/git"
)

func Begin() {
	if !git.IsGitRepo() {
		fmt.Println("Not executed from a git repository")
		return
	}

	localContributors, err := contributors.GetLocalContributors()
	if err != nil {
		fmt.Println(err)
		return
	}

	repoContributors, err := git.GetRepoContributors()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = git.EnablePairingMode(append(localContributors, repoContributors...))
	if err != nil {
		fmt.Println(err)
		return
	}
}
