package main

import (
	"fmt"
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("static"))
	fmt.Printf("%+v", fileServer)
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", app.fileBrowserHandler)
	return mux
}
