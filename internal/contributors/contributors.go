package contributors

import (
	"bufio"
	"fmt"
	"os"

	"github.com/inverse/git-pair/internal/util"
)

const contributorFileName = ".contributors.txt"

func GetLocalContributors() ([]string, error) {
	homeDir := util.GetHomeDir()
	file, err := os.Open(fmt.Sprintf("%s/%s", homeDir, contributorFileName))
	if os.IsNotExist(err) {
		return make([]string, 0), nil
	}

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
