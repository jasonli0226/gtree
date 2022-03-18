package tree

import "github.com/urfave/cli/v2"

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
	Dest string
}

func (gm *Make) Run() {}
