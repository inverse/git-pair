package main

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/inverse/git-pair/internal/contributors"
	"github.com/inverse/git-pair/internal/git"
	"github.com/inverse/git-pair/internal/util"
)

func Start() {
	localContributors, err := contributors.GetLocalContributors()
	if err != nil {
		fmt.Printf("Failed to load local contributors: %s\n", err)
		return
	}

	repoContributors, err := git.GetRepoContributors()
	if err != nil {
		fmt.Printf("Failed to load repo contributors: %s\n", err)
		return
	}

	allContributors := util.UniqueStrings(append(localContributors, repoContributors...))
	if len(allContributors) == 0 {
		fmt.Println("No contributors found")
		return
	}

	selectedContributors := []string{}
	prompt := &survey.MultiSelect{
		Message: "Who's pairing:",
		Options: allContributors,
	}

	err = survey.AskOne(prompt, &selectedContributors)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(selectedContributors) == 0 {
		fmt.Println("You must select at least one contributor")
		return
	}

	err = git.EnablePairingMode(selectedContributors)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("🚀 Pairing mode started")
}
