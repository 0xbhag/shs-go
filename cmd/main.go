package main

import (
	"flag"
	"fmt"
	"log"
	"net"
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
	pathFlag := flag.String("path", ".", "Directory to start http server in.")
	flag.Parse()
	port := ":" + strconv.Itoa(*portFlag)
	dlPath, err := setBrowserPath(pathFlag)
	if err != nil {
		log.Fatal("Input path has problems!")
		os.Exit(1)
	}
	app := application{DownloadPath: dlPath}

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to bind to port %s: %v", port, err)
	}
	fmt.Printf("SHS running on port %s\nShared directory : %s", port, dlPath)

	// log.Fatal(http.ListenAndServe(port, app.routes()))
	log.Fatal(http.Serve(listener, app.routes()))
}
