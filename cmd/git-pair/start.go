package main

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/inverse/git-pair/internal/contributors"
	"github.com/inverse/git-pair/internal/git"
	"github.com/inverse/git-pair/internal/util"
)

func Start() error {
	localContributors, err := contributors.GetLocalContributors()
	if err != nil {
		return fmt.Errorf("failed to load local contributors: %w", err)
	}

	repoContributors, err := git.GetRepoContributors()
	if err != nil {
		return fmt.Errorf("failed to load repo contributors: %w", err)
	}

	allContributors := util.UniqueStrings(append(localContributors, repoContributors...))
	if len(allContributors) == 0 {
		fmt.Println("⚠️ No contributors found")
		return nil
	}

	var selectedContributors []string

	prompt := &survey.MultiSelect{
		Message: "Who's pairing:",
		Options: allContributors,
	}

	if err := survey.AskOne(prompt, &selectedContributors); err != nil {
		return err
	}

	if len(selectedContributors) == 0 {
		fmt.Println("You must select at least one contributor")
		return nil
	}

	if err := git.EnablePairingMode(selectedContributors); err != nil {
		return err
	}

	fmt.Println("🚀 Pairing mode started")
	return nil
}
