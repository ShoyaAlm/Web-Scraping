package models

type UserPreference struct {
	SearchFormat []string `json:"searchFormat"`
	URL          string   `json:"url"`
	Sort         string   `json:"sort"`
	Limit        int      `json:"limit"`
}
