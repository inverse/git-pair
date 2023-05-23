package main

import (
	"fmt"
	"strings"

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
		fmt.Printf("Failed to read template file: %s \n", err)
		return
	}

	for _, contributor := range contributors {
		fmt.Printf("- %s\n", strings.Replace(contributor, git.CoAuthoredBy, "", 1))
	}
}
