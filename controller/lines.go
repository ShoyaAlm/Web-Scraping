package controller

import (
	"fmt"
	"log"
	"webScraper/models"

	"github.com/gocolly/colly"
)

func OnHTMLLinks(collector *colly.Collector, userPref *models.UserPreference) {
	var links []string

	collector.OnHTML("a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		links = append(links, link)
	})

	collector.OnError(func(r *colly.Response, err error) {
		log.Printf("Error scraping page: %v", err)
	})

	collector.OnScraped(func(r *colly.Response) {

		var validLinks []string

		for _, link := range links {
			if link == "" || link == "/" || link == "#" {
				continue
			} else {
				validLinks = append(validLinks, link)
				if len(validLinks) == 4 {
					break
				}
			}

		}

		for _, link := range validLinks {
			fmt.Printf("\n%v\n", link)
		}
	})
}

func OnHTMLParagraphs(collector *colly.Collector, userPref *models.UserPreference) {
	collector.OnHTML("p", func(e *colly.HTMLElement) {
		fmt.Printf("\n%v\n", e.Text)
	})
}

func OnHTMLHeadings(collector *colly.Collector, userPref *models.UserPreference) {
	collector.OnHTML("h1 h2 h3 h4 h5 h6", func(e *colly.HTMLElement) {
		fmt.Printf("\n%v\n", e.Text)
	})
}
