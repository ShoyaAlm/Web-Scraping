package controller

import (
	"fmt"
	"webScraper/models"

	"github.com/gocolly/colly"
)

func OnHTMLVideo(collector *colly.Collector, userPref *models.UserPreference) {
	collector.OnHTML("video", func(e *colly.HTMLElement) {
		fmt.Printf("\n%v\n", e.Attr("src"))
	})
}

func OnHTMLAudio(collector *colly.Collector, userPref *models.UserPreference) {
	collector.OnHTML("audio", func(e *colly.HTMLElement) {
		fmt.Printf("\n%v\n", e.Attr("src"))
	})
}
