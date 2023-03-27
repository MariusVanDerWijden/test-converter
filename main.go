package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
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
		recursive,
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
		in        = c.String(input.Name)
		out       = c.String(output.Name)
		recursive = c.Bool(recursive.Name)
	)
	allFiles, err := os.ReadDir(in)
	if err != nil {
		return err
	}
	if recursive {
		err := filepath.WalkDir(in, func(path string, info fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			abbreviatePath, _ := strings.CutSuffix(path, info.Name())
			if err := handleFile(abbreviatePath, out, info.Name()); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	} else {
		for _, file := range allFiles {
			if err := handleFile(in, out, file.Name()); err != nil {
				return err
			}
		}
	}

	return nil
}

func handleFile(in, out, file string) error {
	ext := path.Ext(file)
	if ext != ".yml" && ext != ".yaml" {
		fmt.Printf("Skipping file %v\n", file)
		return nil
	}
	fmt.Printf("Converting file %v\n", file)
	inputFile := fmt.Sprintf("%v/%v", in, file)
	baseName, _ := strings.CutSuffix(file, ext)
	outputFile := fmt.Sprintf("%v/%v.py", out, baseName)
	if err := convertFile(inputFile, outputFile); err != nil {
		return errors.Join(err, errors.New(outputFile))
	}
	return nil
}
