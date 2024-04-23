package controller

import (
	"fmt"
	"webScraper/models"

	"github.com/gocolly/colly"
)

func OnHTMLLinks(collector *colly.Collector, userPref *models.UserPreference) {
	collector.OnHTML("a", func(e *colly.HTMLElement) {
		if Contains(userPref.SearchFormat, "links") {
			fmt.Printf("\n%v\n", e.Attr("href"))
		}
	})
}

func OnHTMLParagraphs(collector *colly.Collector, userPref *models.UserPreference) {
	collector.OnHTML("p", func(e *colly.HTMLElement) {
		if Contains(userPref.SearchFormat, "paragraphs") {
			fmt.Printf("\n%v\n", e.Text)
		}
	})
}

func OnHTMLHeadings(collector *colly.Collector, userPref *models.UserPreference) {
	collector.OnHTML("h1 h2 h3 h4 h5 h6", func(e *colly.HTMLElement) {
		if Contains(userPref.SearchFormat, "headings") {
			fmt.Printf("\n%v\n", e.Text)
		}
	})
}
