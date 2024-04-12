package models

type PokemonTypeResponse struct {
	DamageRelations struct {
		DoubleDamageFrom []Common `json:"double_damage_from"`
		DoubleDamageTo   []Common `json:"double_damage_to"`
		HalfDamageFrom   []Common `json:"half_damage_from"`
		HalfDamageTo     []Common `json:"half_damage_to"`
		NoDamageFrom     []Common `json:"no_damage_from"`
		NoDamageTo       []Common `json:"no_damage_to"`
	} `json:"damage_relations"`
	GameIndices []struct {
		GameIndex  int64  `json:"game_index"`
		Generation Common `json:"generation"`
	} `json:"game_indices"`
	Generation struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"generation"`
	ID              int      `json:"id"`
	MoveDamageClass Common   `json:"move_damage_class"`
	Moves           []Common `json:"moves"`
	Name            string   `json:"name"`
	Names           []struct {
		Language Common `json:"language"`
		Name     string `json:"name"`
	} `json:"names"`
	PastDamageRelations []interface{} `json:"past_damage_relations"`
	Pokemon             []struct {
		Pokemon Common `json:"pokemon"`
		Slot    int64  `json:"slot"`
	} `json:"pokemon"`
}
