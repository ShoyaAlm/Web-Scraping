package controller

import (
	"fmt"
	"webScraper/models"

	"github.com/gocolly/colly"
)

func OnHTMLVideo(collector *colly.Collector, userPref *models.UserPreference) {
	collector.OnHTML("video", func(e *colly.HTMLElement) {
		if Contains(userPref.SearchFormat, "videos") {
			fmt.Printf("\n%v\n", e.Attr("src"))
		}
	})
}

func OnHTMLAudio(collector *colly.Collector, userPref *models.UserPreference) {
	collector.OnHTML("audio", func(e *colly.HTMLElement) {
		if Contains(userPref.SearchFormat, "audios") {
			fmt.Printf("\n%v\n", e.Attr("src"))
		}
	})
}
