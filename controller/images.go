package controller

import (
	"fmt"
	"webScraper/models"

	"github.com/gocolly/colly"
)

func OnHTMLImages(collector *colly.Collector, userPref *models.UserPreference) {

	var imagesURL []string

	collector.OnHTML("img", func(e *colly.HTMLElement) {
		image := e.Attr("src")
		imagesURL = append(imagesURL, image)
	})

	collector.OnScraped(func(r *colly.Response) {

		var validImages []string
		for _, image := range imagesURL {
			if uniqueImages(image, imagesURL) {
				validImages = append(validImages, image)

			} else if IsInArray(image, repeatedImages) == false {
				fmt.Printf("\n%v\n", image)
				repeatedImages = append(repeatedImages, image)
			}

		}

		for _, image := range validImages {
			fmt.Printf("\n%v\n", image)
		}

	})

}

func OnHTMLImageAlts(collector *colly.Collector, userPref *models.UserPreference) {
	collector.OnHTML("img", func(e *colly.HTMLElement) {

		caption := e.Attr("alt")

		if caption != "" {
			fmt.Printf("\n%v\n", caption)
		}

	})
}

func OnHTMLImageFileTypes(collector *colly.Collector, userPref *models.UserPreference) {
	collector.OnHTML("img", func(e *colly.HTMLElement) {

		src := e.Attr("src")

		fileType := getFileType(src)

		if fileType != "" {
			fmt.Printf("\n%v\n", fileType)
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

var repeatedImages []string

func uniqueImages(imageURL string, images []string) bool {

	var i = 0

	for _, image := range images {
		if image == imageURL {
			i++
		}
	}

	if i == 1 {
		return true
	} else {
		return false
	}

}
