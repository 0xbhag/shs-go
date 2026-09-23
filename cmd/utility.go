package main

import (
	"os"
	"path/filepath"
)

func setBrowserPath(path *string) (string, error) {
	absPath, _ := filepath.Abs(*path)
	_, err := os.Stat(absPath)
	if err != nil {
		return "", nil
	}
	return absPath, nil

}
