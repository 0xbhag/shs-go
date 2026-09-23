package main

import (
	"net/http"
	"os"
	"path/filepath"
	"test/templates"
)

func (app *application) fileBrowserHandler(w http.ResponseWriter, r *http.Request) {
	homePath, err := os.UserHomeDir()
	if err != nil {

	}
	dlPath := filepath.Join(homePath, "Downloads")
	if _, err = os.Stat(dlPath); err != nil {
		os.MkdirAll(dlPath, 0755)
	}
	dirFiles, _ := os.ReadDir(dlPath)
	var files []templates.FileItem

	for _, item := range dirFiles {
		itemInfo, err := item.Info()
		if err != nil {
			return
		}
		files = append(files, templates.FileItem{
			Name: itemInfo.Name(),
			Size: itemInfo.Size(),
		})

	}

	templates.BaseLayout("Home", files).Render(r.Context(), w)

}
