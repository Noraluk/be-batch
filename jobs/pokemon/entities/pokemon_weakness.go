package entities

type PokemonWeakness struct {
	ID            int `gorm:"primaryKey"`
	Name          string
	PokemonID     int
	PokemonTypeID int
}