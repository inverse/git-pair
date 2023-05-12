package main

import (
	"fmt"
	"github.com/inverse/git-pair/internal/diagnostics"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
    app := &cli.App{
        Commands: []*cli.Command{
            {
                Name:    "begin",
                Aliases: []string{"b"},
                Usage:   "begin pairing mode",
                Action: func(*cli.Context) error {
                    fmt.Println("beginning pairing mode")
                    return nil
                },
            },
            {
                Name:    "end",
                Aliases: []string{"e"},
                Usage:   "end pairing mode",
                Action: func(*cli.Context) error {
                    fmt.Println("ending pairing mode")
                    return nil
                },
            },
            {
                Name:    "status",
                Aliases: []string{"s"},
                Usage:   "show pairing status",
                Action: func(*cli.Context) error {
                    fmt.Println("status info")
                    return nil
                },
            },
            {
                Name:    "version",
                Aliases: []string{"v"},
                Usage:   "show version information",
                Action: func(cCtx *cli.Context) error {
                    printVersion()
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
