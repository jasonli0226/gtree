package tree

import (
	"bufio"
	"fmt"
	"gtree/pkg/color"
	"gtree/pkg/utils"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"regexp"

	"github.com/urfave/cli/v2"
)

type SearchDisplayMode uint8

const (
	SearchDisplayNormal SearchDisplayMode = iota
	SearchDisplayFileMode
)

// GetSearchFlags - Get Flags for Search
func GetSearchFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:  "path",
			Usage: "With specified path",
			Value: ".",
		},
		&cli.StringSliceFlag{
			Name:    "ignore-dir",
			Aliases: []string{"I"},
			Usage:   "Ignore specified directory",
		},
		&cli.StringSliceFlag{
			Name:    "pattern",
			Aliases: []string{"p"},
			Usage:   "With specified wildcard",
		},
		&cli.BoolFlag{
			Name:    "file-mode",
			Aliases: []string{"M"},
			Usage:   "Display file mode - directory or file",
		},
		&cli.StringFlag{
			Name:    "target",
			Aliases: []string{"t"},
			Usage:   "with specified target (for file-search mode only)",
			Value:   "",
		},
		&cli.IntFlag{
			Name:    "line",
			Aliases: []string{"l"},
			Usage:   "Number of lines to display (for file-search mode only)",
			Value:   1,
		},
		&cli.BoolFlag{
			Name:    "no-recursive",
			Aliases: []string{"nR"},
			Usage:   "No recursive on searching",
		},
		&cli.BoolFlag{
			Name:    "copy",
			Aliases: []string{"c"},
			Usage:   "Search file in copy mode. Line number will not be displayed",
		},
	}

	return flags
}

// Search - Search Specified Files
type Search struct {
	Target           string
	Mode             SearchDisplayMode
	IsCopy           bool
	NumOfLineDisplay int
	NoRecursive      bool
	IgnoreDirSlice   []string
	PatternSlice     []string

	fileRead            int
	fileWithTargetCount int
	targetCount         int
}

// basicSearch - Basic Search Mode
func (gs *Search) basicSearch(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		path := filepath.Join(path, file.Name())

		isFound := true
		if len(gs.PatternSlice) > 0 && !utils.IsSliceContainsFileMatch(gs.PatternSlice, file.Name()) {
			isFound = false
		}

		if isFound {
			if gs.Mode != SearchDisplayNormal {
				if file.IsDir() {
					fmt.Print(color.Green + "[Directory] \t" + color.Reset)
				} else if file.Mode()&os.ModeSymlink == os.ModeSymlink {
					fmt.Print(color.Green + "[Symlink] \t" + color.Reset)
				} else {
					fmt.Print(color.Green + "[File] \t" + color.Reset)
				}
			}

			fmt.Println(path)
		}

		if file.IsDir() && !gs.NoRecursive && !utils.IsSliceContainsStr(gs.IgnoreDirSlice, file.Name()) {
			gs.basicSearch(path)
		}
	}
}

// fileSearch - File Search Mode
func (gs *Search) fileSearch(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		path := filepath.Join(path, file.Name())

		if !file.IsDir() && !(file.Mode()&os.ModeSymlink == os.ModeSymlink) {
			if len(gs.PatternSlice) > 0 {
				if utils.IsSliceContainsFileMatch(gs.PatternSlice, file.Name()) {
					gs.scanFile(path)
				}
			} else {
				gs.scanFile(path)
			}
			continue
		}

		if gs.NoRecursive {
			continue
		}

		if !utils.IsSliceContainsStr(gs.IgnoreDirSlice, file.Name()) {
			gs.fileSearch(path)
		}
	}
}

// scanFile - Scan a single file
func (gs *Search) scanFile(path string) {
	lineSlice := make([]string, gs.NumOfLineDisplay)
	lastCounterFound := -1
	counter := 0

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	gs.fileRead++
	defer file.Close()

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	for scanner.Scan() {
		line := scanner.Text()
		matched, err := regexp.MatchString(gs.Target, line)
		if err != nil {
			log.Fatal(err)
		}

		counter++

		if matched {
			if lastCounterFound == -1 {
				fmt.Println(color.Yellow + "Reading File - " + path + color.Reset)
			}
			fmt.Println()

			for i, item := range lineSlice {
				if i == counter-1 {
					break
				}
				lineNum := math.Max(float64(counter-gs.NumOfLineDisplay), 1)
				if lastCounterFound == -1 || int(lineNum)+i > lastCounterFound+gs.NumOfLineDisplay {
					gs.displayScanLine(int(lineNum)+i, item, false)
				}
			}

			gs.displayScanLine(counter, line, true)
			lastCounterFound = counter
			gs.targetCount++
		} else if lastCounterFound != -1 && counter-lastCounterFound <= gs.NumOfLineDisplay {
			gs.displayScanLine(counter, line, false)
		}

		if counter <= gs.NumOfLineDisplay {
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
		gs.fileWithTargetCount++
		fmt.Println()
	}
}

func (gs *Search) displayScanLine(lineNum int, line string, isFound bool) {
	lineNumStr := ""
	if !gs.IsCopy {
		lineNumStr = fmt.Sprintf("%d \t", lineNum)
	}
	if isFound {
		fmt.Println(color.Green + lineNumStr + line + color.Reset)
	} else {
		fmt.Println(lineNumStr + line)
	}
}

// Run - Start to Run
func (gs *Search) Run(path string) {
	// add default ignore directory
	gs.IgnoreDirSlice = append(gs.IgnoreDirSlice, "node_modules", ".git")

	if gs.Target != "" {
		gs.fileSearch(path)
		fmt.Print(color.Blue)
		fmt.Println("[Report]")
		fmt.Println("File(s) Read : \t\t", gs.fileRead)
		fmt.Println("File(s) with Target : \t", gs.fileWithTargetCount)
		fmt.Println("Target Line Found : \t", gs.targetCount)
		fmt.Print(color.Reset)
	} else {
		gs.basicSearch(path)
	}
}
