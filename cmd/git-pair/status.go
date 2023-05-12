package main

import (
	"fmt"
	"github.com/inverse/git-pair/internal/git"
)

func Status() {
	fmt.Println("status info")

	fmt.Println(git.IsGitRepo())

}
