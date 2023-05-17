package main

import (
	"fmt"
	"log"
	"os"

	"github.com/inverse/git-pair/internal/diagnostics"
	"github.com/urfave/cli/v2"
)

func main() {
	cli.VersionPrinter = func(cCtx *cli.Context) {
		printVersion()
	}

	app := &cli.App{
		Version: diagnostics.Version,
		Commands: []*cli.Command{
			{
				Name:    "begin",
				Aliases: []string{"b"},
				Usage:   "begin pairing mode",
				Action: func(*cli.Context) error {
					Begin()
					return nil
				},
			},
			{
				Name:    "end",
				Aliases: []string{"e"},
				Usage:   "end pairing mode",
				Action: func(*cli.Context) error {
					End()
					return nil
				},
			},
			{
				Name:    "status",
				Aliases: []string{"s"},
				Usage:   "show pairing status",
				Action: func(*cli.Context) error {
					Status()
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func printVersion() {
	fmt.Printf(
		"version: %s\ncommit: %s\nbuild time: %s\n",
		diagnostics.Version, diagnostics.Commit, diagnostics.BuildDate,
	)
}
