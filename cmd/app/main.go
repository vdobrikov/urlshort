package main

import (
	"flag"
	"fmt"
	"net/http"

	"../../pkg"
)

func main() {
	defaultFilename := "resources/urls.yaml"
	var filename string
	flag.StringVar(&filename, "file", defaultFilename, "yaml/json file with URLs DB")
	flag.StringVar(&filename, "f", defaultFilename, "yaml/json file with URLs DB")

	flag.Parse()

	if filename == "" {
		flag.Usage()
	}

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/golang": "http://golang.org",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yamlHandler, err := urlshort.YAMLHandler(filename, mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	_ = http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/notfound", notFound)
	return mux
}

func notFound(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "Not found")
}

func index(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "URL shortener")
}
