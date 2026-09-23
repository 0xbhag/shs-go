package main

import (
	"os"
	"path/filepath"
)

func setBrowserPath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {

	}
	dlPath := filepath.Join(homePath, "Downloads")
	if _, err = os.Stat(dlPath); err != nil {
		os.MkdirAll(dlPath, 0755)
	}
	return dlPath, nil

}
