package main

import (
	"fmt"

	"github.com/inverse/git-pair/internal/git"
)

func Info() {
	if !git.IsGitRepo() {
		fmt.Println("Not executed from a git repository")
		return
	}

	if !git.PairingModeEnabled() {
		fmt.Println("Pairing mode is not enabled")
		return
	}

	fmt.Println("🔧 Pairing mode is enabled")

	contributors, err := git.ReadTemplateFile()

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, contributor := range contributors {
		fmt.Printf("- %s\n", contributor)
	}
}
