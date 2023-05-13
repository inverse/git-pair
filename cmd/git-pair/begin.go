package main

import (
	"fmt"
	"github.com/inverse/git-pair/internal/git"
)

func Begin() {
	fmt.Println("Beginning pairing mode")
	git.GetRepoContributors()
}
