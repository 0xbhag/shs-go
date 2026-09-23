package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Page struct {
	Title string
	Body  []byte
}
type application struct {
	DownloadPath string
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func load_page(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	portFlag := flag.Int("port", 8080, "HTTP server port")
	flag.Parse()
	port := ":" + strconv.Itoa(*portFlag)
	dlPath, err := setBrowserPath()
	if err != nil {
		return
	}
	app := application{DownloadPath: dlPath}

	log.Fatal(http.ListenAndServe(port, app.routes()))
}
