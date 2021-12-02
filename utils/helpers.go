package utils

import (
	"os"
	"strings"
)

func StringInSlice(str string, slice []string) bool {
	var temp []string
	for _, s := range slice {
		temp = append(temp, strings.ToLower(s))
	}

	for _, s := range temp {
		if s == str {
			return true
		}
	}
	return false
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
