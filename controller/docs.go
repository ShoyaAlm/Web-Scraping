package controller

import (
	"fmt"
	"path/filepath"
	"strings"
	"webScraper/models"

	"github.com/gocolly/colly"
)

func OnHTMLPDFDoc(collector *colly.Collector, userPref *models.UserPreference) {

	collector.OnHTML("a", func(e *colly.HTMLElement) {

		link := e.Attr("href")

		if isDocument(link) {
			extension := getFileExtension(link)

			switch extension {
			case ".pdf":
				if Contains(userPref.SearchFormat, "pdf") {
					printPDF(link)
				}

			case ".doc", "docx":
				if Contains(userPref.SearchFormat, "doc") {
					printDoc(link)
				}

			}

		}

	})

}

func printPDF(link string) {
	fmt.Printf("\n%v\n", link)
}

func printDoc(link string) {
	fmt.Printf("\n%v\n", link)
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

func getFileExtension(link string) string {
	ext := filepath.Ext(link)
	if ext == "" {
		return ""
	}
	return strings.ToLower(ext)
}

// func endsWith(s, suffix string) bool {
// 	return strings.HasSuffix(strings.ToLower(s), suffix)
// }

func OnHTMLTables(collector *colly.Collector, userPref *models.UserPreference) {

	collector.OnHTML("table", func(e *colly.HTMLElement) {
		if Contains(userPref.SearchFormat, "tables") {
			fmt.Printf("\n%v\n", e.Text)
		}
	})
}
