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
		4. Make Directories/Files
		`
	app.Version = "1.1.14"
	app.Commands = []*cli.Command{
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "List all Directories/Files",
			Flags:   tree.GetListFlags(),
			Action: func(c *cli.Context) error {
				isShowFileSize := c.Bool("size")
				isShowTotalSize := c.Bool("total")
				ignoreDirSlice := c.StringSlice("ignore-dir")
				ignoreFileSlice := c.StringSlice("ignore")
				startPath := c.String("path")
				patternSlice := c.StringSlice("pattern")

				var list = tree.List{IsShowFileSize: isShowFileSize, IgnoreDirSlice: ignoreDirSlice,
					PatternSlice: patternSlice, IgnoreFileSlice: ignoreFileSlice}

				list.Run(startPath, isShowTotalSize)
				return nil
			},
		},
		{
			Name:    "remove",
			Aliases: []string{"rm"},
			Usage: `Remove Directories/Files
			Examples:
			1 - gtree rm -t node_modules -R
			2 - gtree rm -p *.ts -R`,
			Flags: tree.GetRemoveFlags(),
			Action: func(c *cli.Context) error {
				startPath := c.String("path")
				isRecursive := c.Bool("Recursive")
				target := c.String("target")
				pattern := c.String("pattern")

				var remove = tree.Remove{IsRecursive: isRecursive, Target: target, Pattern: pattern}

				remove.Run(startPath)
				return nil
			},
		},
		{
			Name:    "scan",
			Aliases: []string{"sc"},
			Usage: `Scan for specified Files. By default, links are not followed.
			Examples:
			1 - gtree sc -I .git -p *.ts
			2 - gtree sc -I .git -p *.go -f -t fun -l 3`,
			Flags: tree.GetSearchFlags(),
			Action: func(c *cli.Context) error {
				startPath := c.String("path")
				target := c.String("target")
				ignoreDirSlice := c.StringSlice("ignore-dir")
				patternSlice := c.StringSlice("pattern")
				mode := tree.SearchDisplayNormal
				numOfLine := int(math.Max(float64(c.Int("line")), 1))
				noRecursive := c.Bool("no-recursive")
				isCopy := c.Bool("copy")

				if c.Bool("file-mode") {
					mode = tree.SearchDisplayFileMode
				}

				var search = tree.Search{Target: target, IgnoreDirSlice: ignoreDirSlice,
					Mode: mode, PatternSlice: patternSlice, IsCopy: isCopy,
					NumOfLineDisplay: numOfLine, NoRecursive: noRecursive}

				search.Run(startPath)
				return nil
			},
		}, {
			Name:    "make",
			Aliases: []string{"mk"},
			Usage: `Make directories or files
			Examples:
			1 - gtree mk -d layer_01/layer_01_01/hello.ts
			2 - gtree mk -d layer_01\layer_01_02\world.ts`,
			Flags: tree.GetMakeFlags(),
			Action: func(c *cli.Context) error {
				destination := c.String("dest")

				var make = tree.Make{}

				make.Run(destination)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
