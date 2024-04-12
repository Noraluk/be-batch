package models

type Ability struct {
	Ability  Common `json:"ability"`
	IsHidden bool   `json:"is_hidden"`
	Slot     int    `json:"slot"`
}
