package tree

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"gtree/pkg/utils"
)

type SearchDisplayMode uint8

const (
	SearchDisplayNormal SearchDisplayMode = iota
	SearchDisplayFileMode
)

// Search - Search specified Files
type Search struct {
	Target           string
	IgnoreDir        string
	Mode             SearchDisplayMode
	Pattern          string
	IsSearchFile     bool
	NumOfLineDisplay int
	color            utils.Color
}

func (s *Search) basicSearch(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		path := filepath.Join(path, file.Name())

		if s.Pattern != "" {
			if match, _ := filepath.Match(s.Pattern, file.Name()); match {
				if s.Mode != SearchDisplayNormal {
					if file.IsDir() {
						fmt.Print(s.color.Green + "[Directory] \t" + s.color.Reset)
					} else {
						fmt.Print(s.color.Green + "[File] \t" + s.color.Reset)
					}
				}
				fmt.Println(path)
			}
		}

		isIgnoreDir := s.IgnoreDir != "" && file.Name() == s.IgnoreDir
		if file.IsDir() && !isIgnoreDir {
			s.Run(path)
		}
	}
}

func (s *Search) fileSearch(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		path := filepath.Join(path, file.Name())

		if s.Pattern != "" {
			if match, _ := filepath.Match(s.Pattern, file.Name()); match {
				s.scanFile(path)
			}
		}

		isIgnoreDir := s.IgnoreDir != "" && file.Name() == s.IgnoreDir
		if file.IsDir() && !isIgnoreDir {
			s.Run(path)
		}
	}
}

func (s *Search) scanFile(path string) {
	lineSlice := make([]string, s.NumOfLineDisplay)
	lastCounterFound := -1
	counter := 0

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		matched, err := regexp.MatchString(s.Target, line)
		if err != nil {
			log.Fatal(err)
		}

		if matched {
			if lastCounterFound == -1 {
				fmt.Println(s.color.Yellow + "Reading File - " + path + s.color.Reset)
			}
			fmt.Println()

			for _, item := range lineSlice {
				fmt.Println(item)
			}

			fmt.Println(s.color.Green + line + s.color.Reset)
			lastCounterFound = counter
		} else if lastCounterFound != -1 && counter-lastCounterFound <= s.NumOfLineDisplay {
			fmt.Println(line)
		}

		counter++
		if counter <= s.NumOfLineDisplay {
			lineSlice[counter-1] = line
		} else {
			lineSlice = lineSlice[1:]
			lineSlice = append(lineSlice, line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if lastCounterFound != -1 {
		fmt.Println()
	}
}

// Run - Start to Run
func (s *Search) Run(path string) {
	s.color = utils.Color{}
	s.color.Init()

	if s.IsSearchFile {
		s.fileSearch(path)
	} else {
		s.basicSearch(path)
	}
}
