package tree

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

// GetRemoveFlags - Get the flags for Remove
func GetRemoveFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "Recursive",
			Aliases: []string{"R"},
			Usage:   "Remove within directories recursively. By default, links are not followed.",
		},
		&cli.StringFlag{
			Name:    "target",
			Aliases: []string{"t"},
			Usage:   "Remove with specified directories/files name",
			Value:   "jason_is_handsome",
		},
		&cli.StringFlag{
			Name:  "path",
			Usage: "Remove with specified path",
			Value: ".",
		},
		&cli.StringFlag{
			Name:    "pattern",
			Aliases: []string{"p"},
			Usage:   "Remove with specified wildcard",
			Value:   "jason_is_handsome",
		},
	}

	return flags
}

// Remove - Remove Folders/Files
type Remove struct {
	IsRecursive bool
	Target      string
	Pattern     string
}

// Run - Run to remove folders/files
func (gr *Remove) Run(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		path := filepath.Join(path, file.Name())
		isTarget := gr.Target != "" && file.Name() == gr.Target
		isPattern, _ := filepath.Match(gr.Pattern, file.Name())

		isSymLink := file.Mode()&os.ModeSymlink == os.ModeSymlink

		if isTarget || isPattern {
			if file.IsDir() {
				os.RemoveAll(path)
				fmt.Println("Deleted directory ======== \t", path)
			} else if !isSymLink && !file.IsDir() {
				os.Remove(path)
				fmt.Println("Deleted file ======== \t", path)
			}
		} else if file.IsDir() && gr.IsRecursive {
			gr.Run(path)
		}
	}
}
