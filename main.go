package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/urfave/cli/v2"
)

var convertCommand = &cli.Command{
	Name:   "convert",
	Usage:  "Converts a directory of files to python tests",
	Action: convertCmd,
	Flags: []cli.Flag{
		input,
		output,
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
	var (
		in  = c.String(input.Name)
		out = c.String(output.Name)
	)
	allFiles, err := ioutil.ReadDir(in)
	if err != nil {
		return err
	}
	for _, file := range allFiles {
		ext := path.Ext(file.Name())
		if ext != ".yml" && ext != ".yaml" {
			fmt.Printf("Skipping file %v\n", file.Name())
			continue
		}
		fmt.Printf("Converting file %v\n", file.Name())
		inputFile := fmt.Sprintf("%v/%v", in, file.Name())
		baseName, _ := strings.CutSuffix(file.Name(), ext)
		outputFile := fmt.Sprintf("%v/%v.py", out, baseName)
		if err := convertFile(inputFile, outputFile); err != nil {
			return errors.Join(err, errors.New(outputFile))
		}
	}
	return nil
}
