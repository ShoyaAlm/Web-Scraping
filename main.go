package main

import (
	"encoding/json"
	"log"
	"net/http"
	_ "os"

	"webScraper/controller"
	"webScraper/models"

	"github.com/gocolly/colly"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./page")))
	http.HandleFunc("/submit", handleSubmit)
	http.ListenAndServe(":3000", nil)
}

func handleSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestData struct {
		URL          string `json:"url"`
		Filter       string `json:"filter"`
		SearchFormat string `json:"searchFormat"`
	}

	// url := "https://webscraper.io/test-sites/e-commerce/allinone"

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Failed to parse from data", http.StatusBadRequest)
		log.Printf("Error encoding request body: %v", err)
		return
	}

	filterSlice := []string{requestData.Filter}

	var searchFormatSlice []string

	if len(requestData.SearchFormat) == 0 {
		searchFormatSlice = []string{requestData.SearchFormat}
	} else {
		searchFormatSlice = append(searchFormatSlice, requestData.SearchFormat)
	}

	var formData = &models.UserPreference{
		URL:          requestData.URL,
		SearchFormat: searchFormatSlice,
		Filter:       filterSlice,
	}

	collector := colly.NewCollector()

	switch formData.SearchFormat[0] {

	case "lines":
		controller.OnHTMLLinks(collector, formData)
		controller.OnHTMLParagraphs(collector, formData)
		controller.OnHTMLHeadings(collector, formData)

	case "images":
		controller.OnHTMLImages(collector, formData)
		controller.OnHTMLImageFileTypes(collector, formData)

	case "webFormat":
		controller.OnHTMLWebFormatJSON(collector, formData)
		controller.OnHTMLWebFormatXML(collector, formData)

	case "webDocs":
		controller.OnHTMLPDFDoc(collector, formData)

	case "multiMedia":
		controller.OnHTMLVideo(collector, formData)
		controller.OnHTMLAudio(collector, formData)

	}

	err := collector.Visit(requestData.URL)
	if err != nil {
		log.Printf("Error visiting URL: %v", err)
		http.Error(w, "Error visiting URL", http.StatusBadRequest)
		return
	}

	log.Printf("Received form data: %+v", formData)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Form data received successfully"))

}

// url := "https://www.jobstreet.vn/j?sp=search&q=C%C3%B4ng+ngh%E1%BB%87+th%C3%B4ng+tin&l="
