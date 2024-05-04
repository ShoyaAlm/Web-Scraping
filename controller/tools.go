package controller

import (
	"sort"
	"webScraper/models"
)

func IsInArray(variable string, infos []string) bool {
	for _, info := range infos {
		if info == variable {
			return true
		}
	}

	return false
}

func LimitedOutput(limit int, infos []string) []string {
	var limitedInfo []string
	i := 0
	if limit != 0 {
		for _, info := range infos {
			limitedInfo = append(limitedInfo, info)
			i++
			if i == limit {
				break
			}
		}
	}

	return limitedInfo
}

func SortInfo(validInfos []string, userPref *models.UserPreference) []string {

	if userPref.Sort == "Low to High" || userPref.Sort == "High to Low" {
		return SortByLength(validInfos, userPref.Sort)
	} else if userPref.Sort == "A-Z order" || userPref.Sort == "Z-A order" {
		return SortByAlphabet(validInfos, userPref.Sort)
	}

	return validInfos

}

func SortByLength(validInfos []string, option string) []string {

	if option == "High to Low" {

		for i := 0; i < len(validInfos)-1; i++ {
			for j := i + 1; j < len(validInfos); j++ {

				if len(validInfos[i]) < len(validInfos[j]) {
					tmp := validInfos[j]
					validInfos[j] = validInfos[i]
					validInfos[i] = tmp
				}

			}
		}

	} else {
		for i := 0; i < len(validInfos)-1; i++ {
			for j := i + 1; j < len(validInfos); j++ {

				if len(validInfos[i]) > len(validInfos[j]) {
					tmp := validInfos[j]
					validInfos[j] = validInfos[i]
					validInfos[i] = tmp
				}

			}
		}
	}

	return validInfos

}

func SortByAlphabet(validInfos []string, option string) []string {

	if option == "A-Z order" {
		sort.Strings(validInfos)
	} else {
		sort.Sort(sort.Reverse(sort.StringSlice(validInfos)))
	}

	return validInfos
}
