package main

import (
	"gtree/pkg/tree"
	"log"
	"math"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "Go Tree"
	app.Usage =
		`Helper Cli for the below features:
		1. List all Directories/Files
		2. Search for specified Files
		3. Remove Directories/Files
		`
	app.Version = "1.1.10"
	app.Commands = []*cli.Command{
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "List all Directories/Files",
			Flags:   tree.GetListFlags(),
			Action: func(c *cli.Context) error {
				showFileSize := c.Bool("size")
				showTotalSize := c.Bool("total")
				ignoreDirSlice := c.StringSlice("ignore-dir")
				ignoreFileSlice := c.StringSlice("ignore")
				path := c.String("path")
				patternSlice := c.StringSlice("pattern")

				var list = tree.List{ShowFileSize: showFileSize, StartPath: path,
					ShowTotalSize: showTotalSize, IgnoreDirSlice: ignoreDirSlice,
					PatternSlice: patternSlice, IgnoreFileSlice: ignoreFileSlice}

				list.Run()
				return nil
			},
		},
		{
			Name:    "remove",
			Aliases: []string{"rm"},
			Usage:   "Remove Directories/Files",
			Flags:   tree.GetRemoveFlags(),
			Action: func(c *cli.Context) error {
				path := c.String("path")
				isRecursive := c.Bool("recursive")
				target := c.String("target")
				pattern := c.String("pattern")

				var remove = tree.Remove{IsRecursive: isRecursive, Target: target, Pattern: pattern}

				remove.Run(path)
				return nil
			},
		},
		{
			Name:    "search",
			Aliases: []string{"sc"},
			Usage:   "Search for specified Files",
			Flags:   tree.GetSearchFlags(),
			Action: func(c *cli.Context) error {
				path := c.String("path")
				target := c.String("target")
				ignoreDirSlice := c.StringSlice("ignore-dir")
				patternSlice := c.StringSlice("pattern")
				mode := tree.SearchDisplayNormal
				isSearchFile := c.Bool("file-search")
				numOfLine := int(math.Max(float64(c.Int("line")), 1))
				noRecursive := c.Bool("no-recursive")

				if c.Bool("file-mode") {
					mode = tree.SearchDisplayFileMode
				}

				var search = tree.Search{Target: target, IgnoreDirSlice: ignoreDirSlice, Mode: mode,
					PatternSlice: patternSlice, IsSearchFile: isSearchFile,
					NumOfLineDisplay: numOfLine, NoRecursive: noRecursive}

				search.Run(path)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
