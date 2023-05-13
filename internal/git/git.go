package git

import (
	"fmt"
	"github.com/inverse/git-pair/internal/util"
	"log"
	"os/exec"
	"strings"
)

const templateFileName string = ".git/template"

func templateFilePath() string {
	out, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()

	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%s/%s", strings.TrimSpace(string(out)), templateFileName)
}

func templateFileExists() bool {
	templatePath := templateFilePath()
	return util.CheckFileExists(templatePath)
}

func IsGitRepo() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	if err := cmd.Run(); err != nil {
		return false
	}

	return true
}

func GetRepoContributors() []string {
	out, err := exec.Command("git", "shortlog", "-e", "-s").Output()
	if err != nil {
		fmt.Println(err)
	}

	contributors := strings.Split(string(out), "\r\n")

	fmt.Println(contributors)

	return contributors
}

func PairingModeEnabled() bool {
	return templateFileExists()
}
