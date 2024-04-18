package models

type PokemonItemsResponse struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []Common `json:"results"`
}

type Common struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
