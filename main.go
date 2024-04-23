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
		SearchFormat: []string{"image-types"},
		URL:          url,
		Filter:       nil,
	}

	controller.OnHTMLImageFileTypes(collector, userPref)

	collector.Visit(url)

	// http.HandleFunc("/", handler)
	// log.Fatal(http.ListenAndServe(":8080", nil))

}
