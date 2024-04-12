package entities

type PokemonAbility struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	PokemonID int
}
