package controller

import (
	"fmt"
	"webScraper/models"

	"github.com/gocolly/colly"
)

func OnHTMLWebFormatJSON(collector *colly.Collector, userPref *models.UserPreference) {
	collector.OnHTML("script", func(e *colly.HTMLElement) {

		if Contains(userPref.SearchFormat, "json") {

			contentType := e.Attr("type")

			if contentType == "application/json" {

				jsonData := e.Text
				fmt.Printf("\n%v\n", jsonData)

			}

		}
	})
}

func OnHTMLWebFormatXML(collector *colly.Collector, userPref *models.UserPreference) {
	collector.OnHTML("script", func(e *colly.HTMLElement) {
		if Contains(userPref.SearchFormat, "xml") {
			contentType := e.Attr("type")

			if contentType == "application/xml" {

				xmlData := e.Text
				fmt.Printf("\n%v\n", xmlData)
			}

		}
	})
}
