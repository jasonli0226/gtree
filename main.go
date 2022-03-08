package main

import (
	"gtree/utils"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "Go Tree"
	app.Usage = "cli tree mocking application"
	app.Version = "1.0.0"
	app.Commands = []cli.Command{
		{
			Name:      "list",
			ShortName: "ls",
			Usage:     "List all directories/files",
			Flags:     getListFlags(),
			Action: func(c *cli.Context) error {
				showFileSize := c.Bool("size")
				showTotalSize := c.Bool("total")
				ignoreFolder := c.String("ignore-folder")
				IgnoreFile := c.String("ignore")
				path := c.String("path")
				pattern := c.String("pattern")

				var list = utils.List{ShowFileSize: showFileSize, StartPath: path,
					ShowTotalSize: showTotalSize, IgnoreFolder: ignoreFolder,
					Pattern: pattern, IgnoreFile: IgnoreFile}

				list.Run()
				return nil
			},
		},
		{
			Name:      "remove",
			ShortName: "rm",
			Usage:     "Remove directories/files",
			Flags:     getRemoveFlags(),
			Action: func(c *cli.Context) error {
				isRecursive := c.Bool("recursive")
				target := c.String("target")
				path := c.String("path")
				pattern := c.String("pattern")

				var remove = utils.Remove{IsRecursive: isRecursive, Target: target, Pattern: pattern}

				remove.Run(path)
				return nil
			},
		},
	}

	app.Run(os.Args)
}

func getListFlags() []cli.Flag {
	flags := []cli.Flag{
		cli.BoolFlag{
			Name:  "size, s",
			Usage: "show file size",
		},
		cli.BoolFlag{
			Name:  "total, t",
			Usage: "show file total size",
		},
		cli.StringFlag{
			Name:  "ignore, i",
			Usage: "list files that are not in the specified patterns",
			Value: "",
		},
		cli.StringFlag{
			Name:  "ignore-folder, I",
			Usage: "list without specified folder",
			Value: "",
		},
		cli.StringFlag{
			Name:  "path",
			Usage: "list specified path",
			Value: ".",
		},
		cli.StringFlag{
			Name:  "pattern, p",
			Usage: "list with specified wildcard",
			Value: "",
		},
	}

	return flags
}

func getRemoveFlags() []cli.Flag {
	flags := []cli.Flag{
		cli.BoolFlag{
			Name:  "recursive, r",
			Usage: "Remove in recursive mode",
		},
		cli.StringFlag{
			Name:  "target, t",
			Usage: "Specified directories/files name",
			Value: "",
		},
		cli.StringFlag{
			Name:  "path",
			Usage: "Specified path",
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
