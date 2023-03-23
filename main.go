package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var convertCommand = &cli.Command{
	Name:   "airdrop",
	Usage:  "Airdrops to a list of accounts",
	Action: convertCmd,
	Flags: []cli.Flag{
		input,
		ouptut,
	},
}

func initApp() *cli.App {
	app := cli.NewApp()
	app.Name = "Test-Converter"
	app.Usage = "Tool for converting ethereum/test fillers into execution-spec-test tests"
	app.Commands = []*cli.Command{
		convertCommand,
	}
	return app
}

var app = initApp()

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func convertCmd(c *cli.Context) error {
	return nil
}
