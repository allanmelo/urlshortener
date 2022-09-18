package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/allanmelo/urlshortener"
)

func main() {
	var fileName string
	
	if len(os.Args) < 2 {
		fileName = "urls.yaml"
	} else {
		fileName = os.Args[1]
	}

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/doc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshortener.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml, err := urlshortener.ReadFile("./main/" + fileName)
	if err != nil {
		panic(err)
	}

	yamlHandler, err := urlshortener.YAMLHandler(yaml, mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}