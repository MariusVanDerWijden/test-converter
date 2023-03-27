package main

import (
	"github.com/urfave/cli/v2"
)

var (
	input = &cli.StringFlag{
		Name:     "in",
		Usage:    "Specify the input directory",
		Required: true,
	}

	recursive = &cli.BoolFlag{
		Name:     "rec",
		Usage:    "Specify if all subdirectories should be converted as well",
		Required: false,
	}

	output = &cli.StringFlag{
		Name:     "out",
		Usage:    "Specify the output directory",
		Required: true,
	}
)
