package controller

import (
	"fmt"
	"webScraper/models"

	"github.com/gocolly/colly"
)

func OnHTMLImages(collector *colly.Collector, userPref *models.UserPreference) {
	collector.OnHTML("img", func(e *colly.HTMLElement) {
		if Contains(userPref.SearchFormat, "images") {
			fmt.Printf("\n%v\n", e.Attr("src"))
		}
	})
}

func OnHTMLImageAlts(collector *colly.Collector, userPref *models.UserPreference) {
	collector.OnHTML("img", func(e *colly.HTMLElement) {
		if Contains(userPref.SearchFormat, "image-alts") {

			caption := e.Attr("alt")

			if caption != "" {
				fmt.Printf("\n%v\n", caption)
			}

		}
	})
}

func OnHTMLImageFileTypes(collector *colly.Collector, userPref *models.UserPreference) {
	collector.OnHTML("img", func(e *colly.HTMLElement) {
		if Contains(userPref.SearchFormat, "image-types") {

			src := e.Attr("src")

			fileType := getFileType(src)

			if fileType != "" {
				fmt.Printf("\n%v\n", fileType)
			}

		}
	})
}

func getFileType(src string) string {

	if endsWith(src, ".jpg") || endsWith(src, ".jpeg") {
		return "JPEG"
	}

	if endsWith(src, ".svg") {
		return "SVG"
	}

	if endsWith(src, ".png") {
		return "PNG"
	}

	if endsWith(src, ".gif") {
		return "GIF"
	}

	return ""

}

func endsWith(str, suffix string) bool {
	return len(str) >= len(suffix) && str[len(str)-len(suffix):] == suffix
}
