package controller

import (
	"fmt"
	"webScraper/models"

	"github.com/gocolly/colly"
)

func OnHTMLPDFDoc(collector *colly.Collector, userPref *models.UserPreference) {

	collector.OnHTML("a", func(e *colly.HTMLElement) {

		if Contains(userPref.SearchFormat, "documents") {
			link := e.Attr("href")
			if isDocument(link) {
				fmt.Printf("\n%v\n", link)
			}
		}

	})

}

func isDocument(link string) bool {

	documentExtensions := []string{".pdf", ".docx", "doc"}

	for _, ext := range documentExtensions {
		if endsWith(link, ext) {
			return true
		}
	}

	return false
}

func OnHTMLTables(collector *colly.Collector, userPref *models.UserPreference) {

	collector.OnHTML("table", func(e *colly.HTMLElement) {
		if Contains(userPref.SearchFormat, "tables") {
			fmt.Printf("\n%v\n", e.Text)
		}
	})
}
