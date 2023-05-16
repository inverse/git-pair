package git

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/inverse/git-pair/internal/util"
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

func getCurrentBranch() string {
	out, err := exec.Command("git", "branch", "--show-current").Output()

	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(string(out))
}

func IsGitRepo() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	if err := cmd.Run(); err != nil {
		return false
	}

	return true
}

func GetRepoContributors() ([]string, error) {
	out, err := exec.Command("git", "shortlog", "-e", "-s", getCurrentBranch()).Output()
	if err != nil {
		return nil, err
	}

	contributors := strings.Split(string(out), "\r\n")
	for index, contributor := range contributors {
		contributors[index] = strings.TrimSpace(contributor[strings.IndexByte(contributor, '\t'):])
	}

	return contributors, nil
}

func PairingModeEnabled() bool {
	return templateFileExists()
}

func DisablePairingMode() {
	templateFilePath := templateFilePath()
	os.Remove(templateFilePath)
}

func EnablePairingMode(contributors []string) {
	templateFilePath := templateFilePath()

	templateFile, err := os.Create(templateFilePath)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer templateFile.Close()

	for _, contributor := range contributors {
		_, err := templateFile.WriteString(contributor + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func ReadTemplateFile() ([]string, error) {
	templateFilePath := templateFilePath()

	file, err := os.Open(templateFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
