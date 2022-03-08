package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"
)

// GList struct definition
type List struct {
	ShowFileSize  bool
	ShowTotalSize bool
	StartPath     string
	Pattern       string
	IgnoreFolder  string
	IgnoreFile    string
}

// displayFileSize - convert and return the fileSize
func (gl *List) displayFileSize(size int64) string {
	var m float64 = 1024
	var result string
	if size < 1024 {
		return fmt.Sprintf("[%.0f B]", float64(size))
	}
	kb := float64(size) / m
	if kb < 1024 {
		result = fmt.Sprintf("[%.2f KB]", kb)
	} else {
		result = fmt.Sprintf("[%.2f MB]", kb/m)
	}
	return result
}

// ListAllFiles - function to dispaly all the file paths
func (gl *List) listAllFiles(prefixPad string, path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	padding := prefixPad
	for idx, file := range files {
		if file.IsDir() {
			fmt.Println(padding+"|---", file.Name())
			if file.Name() != gl.IgnoreFolder {
				gl.listAllFiles(padding+"|   ", filepath.Join(path, file.Name()))
			}

		} else {
			if gl.Pattern != "" {
				if match, _ := filepath.Match(gl.Pattern, file.Name()); !match {
					continue
				}
			}

			if gl.IgnoreFile != "" {
				if match, _ := filepath.Match(gl.IgnoreFile, file.Name()); match {
					continue
				}
			}

			var paddingFinal string
			if idx == len(files)-1 {
				paddingFinal = padding + "\\---"
			} else {
				paddingFinal = padding + "|---"
			}

			if gl.ShowFileSize {
				fmt.Println(paddingFinal, file.Name(), gl.displayFileSize(file.Size()))
			} else {
				fmt.Println(paddingFinal, file.Name())
			}
		}
	}
}

// showTotalSize - Show total size of the files
func (gl *List) showTotalSize() {
	ch := make(chan int)
	ans := make(chan int64)
	var fileNum int64 = 0

	go gl.sumFileSize(ch, ans)

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
			if gl.Pattern != "" {
				if match, _ := filepath.Match(gl.Pattern, file.Name()); !match {
					continue
				}
			}

			if gl.IgnoreFile != "" {
				if match, _ := filepath.Match(gl.IgnoreFile, file.Name()); match {
					continue
				}
			}

			sum += file.Size()
			ch <- 1
			time.Sleep(time.Millisecond * 50)
		}
	}
	return sum
}

// sumFileSize - sum the file size
func (gl *List) sumFileSize(ch chan int, ans chan int64) {
	sum := gl.loopFileSize(gl.StartPath, ch)
	ans <- sum
	close(ch)
	close(ans)
}

// Run - Run the proccess
func (gl *List) Run() {
	if !gl.ShowTotalSize {
		gl.listAllFiles("  ", gl.StartPath)
		return
	}
	gl.showTotalSize()
}
