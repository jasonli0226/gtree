package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

type SearchMode uint8

const (
	SearchDisplayNormal SearchMode = iota
	SearchDisplayBasic
)

// Search - Search specified Files
type Search struct {
	Target    string
	IgnoreDir string
	Mode      SearchMode
}

// Run - Start to Run
func (s *Search) Run(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		path := filepath.Join(path, file.Name())
		isTarget := s.Target != "" && file.Name() == s.Target

		if isTarget {
			if s.Mode != SearchDisplayNormal {
				if file.IsDir() {
					fmt.Print("[directory] \t")
				} else {
					fmt.Print("[file] \t")
				}
			}
			fmt.Println(path)
		}

		isIgnoreDir := s.IgnoreDir != "" && file.Name() == s.IgnoreDir
		if file.IsDir() && !isIgnoreDir {
			s.Run(path)
		}
	}
}
