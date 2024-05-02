package models

type UserPreference struct {
	SearchFormat []string `json:"searchFormat"`
	URL          string   `json:"url"`
	Filter       []string `json:"filter"`
	Limit        int      `json:"limit"`
}
