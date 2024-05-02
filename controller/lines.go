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
				if isInArray(link, repeatedInfo) {
					fmt.Printf("\n%v\n", link)
				}
				if uniqueLinks(link, links) {
					validLinks = append(validLinks, link)
				} else {
					repeatedInfo = append(repeatedInfo, link)
				}
			}

		}

		for _, link := range validLinks {
			fmt.Printf("\n%v\n", link)
		}

	})

}

func OnHTMLParagraphs(collector *colly.Collector, userPref *models.UserPreference) {

	var pars []string

	collector.OnHTML("p", func(e *colly.HTMLElement) {
		par := e.Text
		pars = append(pars, par) // appending paragraphs
	})

	collector.OnScraped(func(r *colly.Response) {
		var validPars []string

		for _, par := range pars {
			validPars = append(validPars, par)
		}

		for _, par := range validPars {
			fmt.Printf("\n%v\n", par)
		}

	})
}

func OnHTMLHeadings(collector *colly.Collector, userPref *models.UserPreference) {
	collector.OnHTML("h1 h2 h3 h4 h5 h6", func(e *colly.HTMLElement) {
		fmt.Printf("\n%v\n", e.Text)
	})
}

func uniqueLinks(variable string, links []string) bool {
	var i = 0
	var isUnique bool
	for _, link := range links {
		if link == variable {
			i++
		}
	}

	if i == 1 {
		isUnique = true
	} else {
		isUnique = false
	}

	return isUnique
}

var repeatedInfo []string

func isInArray(variable string, infos []string) bool {
	for _, info := range infos {
		if info == variable {
			return true
		}
	}

	return false
}
