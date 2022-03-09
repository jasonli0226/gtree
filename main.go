package main

import (
	"gtree/utils"
	"os"

	"gopkg.in/urfave/cli.v1"
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
	app.Version = "1.0.1"
	app.Commands = []cli.Command{
		{
			Name:      "list",
			ShortName: "ls",
			Usage:     "List all Directories/Files",
			Flags:     getListFlags(),
			Action: func(c *cli.Context) error {
				showFileSize := c.Bool("size")
				showTotalSize := c.Bool("total")
				ignoreDir := c.String("ignore-dir")
				ignoreFile := c.String("ignore")
				path := c.String("path")
				pattern := c.String("pattern")

				var list = utils.List{ShowFileSize: showFileSize, StartPath: path,
					ShowTotalSize: showTotalSize, IgnoreDir: ignoreDir,
					Pattern: pattern, IgnoreFile: ignoreFile}

				list.Run()
				return nil
			},
		},
		{
			Name:      "remove",
			ShortName: "rm",
			Usage:     "Remove Directories/Files",
			Flags:     getRemoveFlags(),
			Action: func(c *cli.Context) error {
				path := c.String("path")
				isRecursive := c.Bool("recursive")
				target := c.String("target")
				pattern := c.String("pattern")

				var remove = utils.Remove{IsRecursive: isRecursive, Target: target, Pattern: pattern}

				remove.Run(path)
				return nil
			},
		},
		{
			Name:      "search",
			ShortName: "sc",
			Usage:     "Search for specified Files",
			Flags:     getSearchFlags(),
			Action: func(c *cli.Context) error {
				path := c.String("path")
				target := c.String("target")
				ignoreDir := c.String("ignore-dir")
				mode := c.Int("mode")

				var search = utils.Search{Target: target, IgnoreDir: ignoreDir, Mode: utils.SearchMode(mode)}

				search.Run(path)
				return nil
			},
		},
	}

	app.Run(os.Args)
}

func getSearchFlags() []cli.Flag {
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "path",
			Usage: "with specified path",
			Value: ".",
		},
		cli.StringFlag{
			Name:  "target, t",
			Usage: "with specified target (file)",
			Value: "jason_is_handsome",
		},
		cli.StringFlag{
			Name:  "ignore-dir, I",
			Usage: "ignore specified directory",
			Value: "",
		},
		cli.IntFlag{
			Name:  "mode, m",
			Usage: "display mode: Normal - 0, Basic - 1",
			Value: int(utils.SearchDisplayNormal),
		},
	}

	return flags
}

func getListFlags() []cli.Flag {
	flags := []cli.Flag{
		cli.BoolFlag{
			Name:  "size, s",
			Usage: "display file size",
		},
		cli.BoolFlag{
			Name:  "total, t",
			Usage: "display file total size",
		},
		cli.StringFlag{
			Name:  "ignore, i",
			Usage: "ignore specified patterns",
			Value: "",
		},
		cli.StringFlag{
			Name:  "ignore-dir, I",
			Usage: "ignore specified directory",
			Value: "",
		},
		cli.StringFlag{
			Name:  "path",
			Usage: "with specified path",
			Value: ".",
		},
		cli.StringFlag{
			Name:  "pattern, p",
			Usage: "with specified wildcard",
			Value: "",
		},
	}

	return flags
}

func getRemoveFlags() []cli.Flag {
	flags := []cli.Flag{
		cli.BoolFlag{
			Name:  "recursive, r",
			Usage: "in recursive mode",
		},
		cli.StringFlag{
			Name:  "target, t",
			Usage: "with specified directories/files name",
			Value: "jason_is_handsome",
		},
		cli.StringFlag{
			Name:  "path",
			Usage: "with specified path",
			Value: ".",
		},
		cli.StringFlag{
			Name:  "pattern, p",
			Usage: "remove with specified wildcard",
			Value: "",
		},
	}

	return flags
}
