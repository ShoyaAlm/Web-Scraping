package main

import (
	"fmt"
	_ "fmt"
	_ "log"
	"net/http"
	_ "net/http"
	_ "os"

	"webScraper/controller"

	"webScraper/models"

	"github.com/gocolly/colly"
)

// type userPreference struct {
// 	searchFormat []string
// 	URL          string
// 	filter       []interface{}
// }

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {

	url := "https://webscraper.io/test-sites/e-commerce/allinone"

	// url := "https://www.jobstreet.vn/j?sp=search&q=C%C3%B4ng+ngh%E1%BB%87+th%C3%B4ng+tin&l="
	var collector = colly.NewCollector()

	collector.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting %v\n", r.URL)
	})

	collector.OnResponse(func(r *colly.Response) {
		fmt.Printf("Got a response from %v\n", r.Request.URL)
	})

	collector.OnError(func(r *colly.Response, e error) {
		fmt.Printf("An error occurred: %v\n", e)
	})

	userPref := &models.UserPreference{
		SearchFormat: []string{"links", "paragraphs"},
		URL:          url,
		Filter:       nil,
	}

	// controller.OnHTMLWebFormatJSON(collector, userPref)

	controller.OnHTMLTables(collector, userPref)

	controller.OnHTMLPDFDoc(collector, userPref)
	controller.OnHTMLParagraphs(collector, userPref)
	controller.OnHTMLLinks(collector, userPref)

	collector.Visit(url)

	// http.HandleFunc("/", handler)
	// log.Fatal(http.ListenAndServe(":8080", nil))

}
