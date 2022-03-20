package tree

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/urfave/cli/v2"
)

func GetMakeFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:    "dest",
			Aliases: []string{"d"},
			Usage:   "destination",
			Value:   "",
		},
	}

	return flags
}

type Make struct {
}

// Run - Run the Make Command
func (gm *Make) Run(target string) {
	if target == "" {
		return
	}

	re := regexp.MustCompile(`[/\\]`)
	destSlice := re.Split(target, -1)
	path := "."

	for _, target := range destSlice {
		path = filepath.Join(path, target)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			var createErr error
			if strings.ContainsAny(target, ".") {
				_, createErr = os.Create(path)
				fmt.Println("Created file \t\t ========", path)
			} else {
				createErr = os.Mkdir(path, 0755)
				fmt.Println("Created directory \t ========", path)
			}
			if createErr != nil {
				log.Fatal(createErr)
			}
		}
	}
}
