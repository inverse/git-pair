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
	templateFilePath := templateFilePath()
	return util.CheckFileExists(templateFilePath)
}

func getCurrentBranch() string {
	out, err := exec.Command("git", "branch", "--show-current").Output()

	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(string(out))
}

func createTemplateFile(contributors []string, templateFilePath string) error {
	templateFile, err := os.Create(templateFilePath)
	if err != nil {
		return err
	}

	defer templateFile.Close()

	for _, contributor := range contributors {
		_, err := templateFile.WriteString(fmt.Sprintf("Co-authored-by: %s \n", contributor))
		if err != nil {
			return err
		}
	}

	return nil
}

func enableGitTemplateConfig(templateFilePath string) error {
	_, err := exec.Command("git", "config", "commit.template", templateFilePath).Output()

	return err
}

func disableGitTemplateConfig() error {
	_, err := exec.Command("git", "config", "--unset", "commit.template").Output()

	return err
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

	toRemove := []int{}
	contributors := strings.Split(string(out), "\n")
	for index, contributor := range contributors {
		tabIndex := strings.IndexByte(contributor, '\t')
		if tabIndex == -1 {
			toRemove = append(toRemove, index)
			continue
		}
		contributors[index] = strings.TrimSpace(contributor[tabIndex:])
	}

	for _, index := range toRemove {
		contributors = append(contributors[:index], contributors[index+1:]...)
	}

	return contributors, nil
}

func PairingModeEnabled() bool {
	return templateFileExists()
}

func DisablePairingMode() error {
	templateFilePath := templateFilePath()
	if !util.CheckFileExists(templateFilePath) {
		return nil
	}

	os.Remove(templateFilePath)

	return disableGitTemplateConfig()
}

func EnablePairingMode(contributors []string) error {
	templateFilePath := templateFilePath()
	err := createTemplateFile(contributors, templateFilePath)
	if err != nil {
		return err
	}

	return enableGitTemplateConfig(templateFilePath)
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
