package utils

import "path/filepath"

// IsSliceContainsStr - Check if the target is in the given slice
func IsSliceContainsStr(strSlice []string, target string) bool {
	for _, s := range strSlice {
		if s == target {
			return true
		}
	}
	return false
}

// IsSliceContainsFileMatch - Check if the filename is matched in the given slice
func IsSliceContainsFileMatch(patternSlice []string, filename string) bool {
	for _, s := range patternSlice {
		if match, _ := filepath.Match(s, filename); match {
			return true
		}
	}
	return false
}
