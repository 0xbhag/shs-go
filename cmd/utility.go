package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func setBrowserPath(path *string) (string, error) {
	absPath, _ := filepath.Abs(*path)
	_, err := os.Stat(absPath)
	if err != nil {
		fmt.Printf("%+v\n", absPath)
		return "", err
	}
	return absPath, nil

}
