package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"shs/templates"
)

func (app *application) fileBrowserHandler(w http.ResponseWriter, r *http.Request) {

	dirFiles, _ := os.ReadDir(app.DownloadPath)
	var files []templates.FileItem

	for _, item := range dirFiles {
		itemInfo, err := item.Info()
		if err != nil {
			return
		}
		if !itemInfo.IsDir() {
			files = append(files, templates.FileItem{
				Name: itemInfo.Name(),
				Size: itemInfo.Size(),
			})
		}

	}

	templates.BaseLayout("Home", files).Render(r.Context(), w)

}

func (app *application) fileDownloadHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.PathValue("filename")
	safe_filename := filepath.Base(filename)
	targetFilePath := filepath.Join(app.DownloadPath, safe_filename)
	if _, err := os.Stat(targetFilePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", safe_filename))
	http.ServeFile(w, r, targetFilePath)
}
