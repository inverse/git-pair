package git

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const templateFileName string = ".git/template"

func checkFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	return !errors.Is(error, os.ErrNotExist)
}

func templateFileExists() bool {
	out, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()

	if err != nil {
		log.Fatal(err)
	}

	templatePath := fmt.Sprintf("%s/%s", strings.TrimSpace(string(out)), templateFileName)
	return checkFileExists(templatePath)
}

func IsGitRepo() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func PairingModeEnabled() bool {
	return templateFileExists()
}
