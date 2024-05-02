package controller

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
