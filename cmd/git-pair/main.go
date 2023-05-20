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
				Name:    "start",
				Aliases: []string{"s"},
				Usage:   "start pairing mode",
				Action: func(*cli.Context) error {
					Start()
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
				Name:    "info",
				Aliases: []string{"i"},
				Usage:   "show pairing info",
				Action: func(*cli.Context) error {
					Info()
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
