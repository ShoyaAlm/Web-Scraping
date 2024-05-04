package controller

import "webScraper/models"

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

	if userPref.Sort == "length" {
		return SortByLength(validInfos)
	}

	return validInfos

}

// sorting descendingly
func SortByLength(validInfos []string) []string {

	for i := 0; i < len(validInfos)-1; i++ {
		for j := i + 1; j < len(validInfos); j++ {
			if len(validInfos[i]) < len(validInfos[j]) {
				tmp := validInfos[j]
				validInfos[j] = validInfos[i]
				validInfos[i] = tmp
			}
		}
	}

	return validInfos

}
