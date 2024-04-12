package entities

type PokemonType struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	PokemonID int
}
