package git

import (
	"log"
	"os/exec"
)

func IsGitRepo() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}
