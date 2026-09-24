package main

import (
	"net/http"
	"shs/static"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	fileServer := http.FileServerFS(static.Files)
	// fmt.Printf("%+v", fileServer)
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", app.fileBrowserHandler)
	mux.HandleFunc("GET /download/{filename}", app.fileDownloadHandler)
	return mux
}
