package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Remove - Remove Folders/Files
type Remove struct {
	IsRecursive bool
	Target      string
	Pattern     string
}

// Run - Run to remove folders/files
func (r *Remove) Run(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		path := filepath.Join(path, file.Name())
		isTarget := r.Target != "" && file.Name() == r.Target
		isPattern, _ := filepath.Match(r.Pattern, file.Name())

		if isTarget || isPattern {
			if file.IsDir() {
				os.RemoveAll(path)
				fmt.Println("deleted directory ==== \t\t", path)
			} else if !file.IsDir() {
				os.Remove(path)
				fmt.Println("deleted file ==== \t\t", path)
			}
		} else {
			if file.IsDir() && r.IsRecursive {
				r.Run(path)
			}
		}
	}
}
