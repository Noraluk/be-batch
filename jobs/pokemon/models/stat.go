package models

type Stat struct {
	BaseStat int    `json:"base_stat"`
	Effort   int    `json:"effort"`
	Stat     Common `json:"stat"`
}
