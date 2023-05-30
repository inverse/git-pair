package main

import (
	"fmt"

	"github.com/inverse/git-pair/internal/git"
)

func End() {
	err := git.DisablePairingMode()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("👋 Finished pairing mode")
}
