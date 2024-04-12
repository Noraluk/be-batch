package models

type PokemonResponse struct {
	Abilities              []Ability     `json:"abilities"`
	BaseExperience         int           `json:"base_experience"`
	Cries                  Cries         `json:"cries"`
	Forms                  []Common      `json:"forms"`
	GameIndices            []GameIndice  `json:"game_indices"`
	Height                 float64       `json:"height"`
	HeldItems              []interface{} `json:"held_items"`
	ID                     int           `json:"id"`
	IsDefault              bool          `json:"is_default"`
	LocationAreaEncounters string        `json:"location_area_encounters"`
	Moves                  []Move        `json:"moves"`
	Name                   string        `json:"name"`
	Order                  int           `json:"order"`
	PastAbilities          []interface{} `json:"past_abilities"`
	PastTypes              []interface{} `json:"past_types"`
	Species                Common        `json:"species"`
	Sprites                Sprites       `json:"sprites"`
	Stats                  []Stat        `json:"stats"`
	Types                  []Type        `json:"types"`
	Weight                 float64       `json:"weight"`
}
