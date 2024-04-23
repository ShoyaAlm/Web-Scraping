package main

import (
	"fmt"
	"log"
	"net/http"
)

type Data struct {
	lines      []interface{} // could be paragraphs, headings or links
	photos     []interface{} // pics & stuff
	webFormat  []interface{} // JSON or XML format
	webDocs    []interface{} // PDFs, Docs, Word files
	multiMedia []interface{} // videos, audio

}

type userPreference struct {
	searchFormat []interface{}
	URL          string
	filter       []interface{}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {

	userPr := userPreference{}

	userPr.URL = "https://www.zoomit.ir"

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
