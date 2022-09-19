package tree

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"

	"github.com/jasonli0226/gtree/pkg/color"
	"github.com/jasonli0226/gtree/pkg/utils"
	"github.com/urfave/cli/v2"
)

// GetListFlags - Get flags for List
func GetListFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "size",
			Aliases: []string{"S"},
			Usage:   "display file size",
		},
		&cli.BoolFlag{
			Name:    "total",
			Aliases: []string{"T"},
			Usage:   "display file total size",
		},
		&cli.StringFlag{
			Name:  "path",
			Usage: "with specified path",
			Value: ".",
		},
		&cli.StringSliceFlag{
			Name:    "pattern",
			Aliases: []string{"p"},
			Usage:   "with specified wildcard",
		},
		&cli.StringSliceFlag{
			Name:    "ignore",
			Aliases: []string{"i"},
			Usage:   "ignore specified patterns",
		},
		&cli.StringSliceFlag{
			Name:    "ignore-dir",
			Aliases: []string{"I"},
			Usage:   "ignore specified directory",
		},
	}

	return flags
}

// GList Struct Definition
type List struct {
	IsShowFileSize  bool
	PatternSlice    []string
	IgnoreFileSlice []string
	IgnoreDirSlice  []string
}

// displayFileSize - Convert and return the fileSize
func (gl *List) displayFileSize(size int64) string {
	var m float64 = 1024
	var result string
	if size < 1024 {
		return color.Green + fmt.Sprintf("[%.0f B]", float64(size)) + color.Reset
	}

	kb := float64(size) / m
	if kb < 1024 {
		result = fmt.Sprintf("[%.2f KB]", kb)
	} else {
		result = fmt.Sprintf("[%.2f MB]", kb/m)
	}

	return color.Green + result + color.Reset
}

// ListAllFiles - Function to dispaly all the file paths
func (gl *List) listAllFiles(prefixPad string, path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	padding := prefixPad
	for idx, file := range files {
		if file.IsDir() {
			fmt.Println(padding+"|---", file.Name())
			if !utils.IsSliceContainsStr(gl.IgnoreDirSlice, file.Name()) {
				gl.listAllFiles(padding+"|   ", filepath.Join(path, file.Name()))
			}

		} else {
			if len(gl.PatternSlice) > 0 && !utils.IsSliceContainsFileMatch(gl.PatternSlice, file.Name()) {
				continue
			}

			if utils.IsSliceContainsFileMatch(gl.IgnoreFileSlice, file.Name()) {
				continue
			}

			var paddingFinal string
			if idx == len(files)-1 {
				paddingFinal = padding + "\\---"
			} else {
				paddingFinal = padding + "|---"
			}

			if gl.IsShowFileSize {
				fmt.Println(paddingFinal, file.Name(), gl.displayFileSize(file.Size()))
			} else {
				fmt.Println(paddingFinal, file.Name())
			}
		}
	}
}

// showTotalSize - Show total size of the files
func (gl *List) showTotalSize(startPath string) {
	ch := make(chan int)
	ans := make(chan int64)
	var fileNum int64 = 0

	go gl.sumFileSize(startPath, ch, ans)

	for {
		select {
		case _, ok := <-ch:
			{
				if !ok {
					break
				}
				fileNum++
				fmt.Print("=")
			}
		case sum, ok := <-ans:
			{
				if !ok {
					return
				}
				fmt.Println("|")
				fmt.Println("Total File Num: \t", fileNum)
				fmt.Println("Total File Size: \t", gl.displayFileSize(sum))
			}
		}
	}
}

// loopFileSize
func (gl *List) loopFileSize(path string, ch chan int) int64 {
	var sum int64 = 0
	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			sum += gl.loopFileSize(filepath.Join(path, file.Name()), ch)
		} else {
			if len(gl.PatternSlice) > 0 && !utils.IsSliceContainsFileMatch(gl.PatternSlice, file.Name()) {
				continue
			}

			if utils.IsSliceContainsFileMatch(gl.IgnoreFileSlice, file.Name()) {
				continue
			}

			sum += file.Size()
			ch <- 1
			time.Sleep(time.Millisecond * 50)
		}
	}

	return sum
}

// sumFileSize - Sum the file size
func (gl *List) sumFileSize(startPath string, ch chan int, ans chan int64) {
	sum := gl.loopFileSize(startPath, ch)
	ans <- sum
	close(ch)
	close(ans)
}

// Run - Run the process
func (gl *List) Run(startPath string, isShowTotalSize bool) {
	// add default ignore directory
	gl.IgnoreDirSlice = append(gl.IgnoreDirSlice, "node_modules", ".git")

	if isShowTotalSize {
		gl.showTotalSize(startPath)
		return
	}

	gl.listAllFiles("  ", startPath)
}
