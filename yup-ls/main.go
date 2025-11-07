package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/ls"
)

const (
	flagLongFormat    = "long"
	flagAllFiles      = "all"
	flagHumanReadable = "human-readable"
	flagRecursive     = "recursive"
	flagReverse       = "reverse"
	flagSortBy        = "sort"
)

func main() {
	app := &cli.App{
		Name:  "ls",
		Usage: "list directory contents",
		UsageText: `ls [OPTIONS] [FILE...]

   List information about the FILEs (the current directory by default).`,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    flagLongFormat,
				Aliases: []string{"l"},
				Usage:   "use a long listing format",
			},
			&cli.BoolFlag{
				Name:    flagAllFiles,
				Aliases: []string{"a"},
				Usage:   "do not ignore entries starting with .",
			},
		&cli.BoolFlag{
			Name:  flagHumanReadable,
			Usage: "with -l, print sizes in human readable format",
		},
			&cli.BoolFlag{
				Name:    flagRecursive,
				Aliases: []string{"R"},
				Usage:   "list subdirectories recursively",
			},
			&cli.BoolFlag{
				Name:    flagReverse,
				Aliases: []string{"r"},
				Usage:   "reverse order while sorting",
			},
			&cli.StringFlag{
				Name:  flagSortBy,
				Usage: "sort by WORD (name, time, size)",
				Value: "name",
			},
		},
		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "ls: %v\n", err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	var params []any

	// Add file/directory arguments
	for i := 0; i < c.NArg(); i++ {
		params = append(params, c.Args().Get(i))
	}

	// Add flags based on CLI options
	if c.Bool(flagLongFormat) {
		params = append(params, LongFormat)
	}
	if c.Bool(flagAllFiles) {
		params = append(params, AllFiles)
	}
	if c.Bool(flagHumanReadable) {
		params = append(params, HumanReadable)
	}
	if c.Bool(flagRecursive) {
		params = append(params, Recursive)
	}
	if c.Bool(flagReverse) {
		params = append(params, Reverse)
	}
	if c.IsSet(flagSortBy) {
		params = append(params, SortBy(c.String(flagSortBy)))
	}

	// Create and execute the ls command
	cmd := Ls(params...)
	return yup.Run(cmd)
}
