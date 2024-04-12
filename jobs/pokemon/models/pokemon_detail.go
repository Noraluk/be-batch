package models

import "be-batch/jobs/pokemon/entities"

type PokemonDetail struct {
	Pokemon           entities.Pokemon
	PokemonAbilities  []entities.PokemonAbility
	PokemonStats      []entities.PokemonStat
	PokemonTypes      []entities.PokemonType
	PokemonWeaknesses []entities.PokemonWeakness
}
