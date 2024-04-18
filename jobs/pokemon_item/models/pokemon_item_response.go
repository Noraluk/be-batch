package models

type PokemonItemResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Cost    int    `json:"cost"`
	Sprites struct {
		Default string `json:"default"`
	} `json:"sprites"`
}
