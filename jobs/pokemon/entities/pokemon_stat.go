package entities

type PokemonStat struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	BaseStat  int
	PokemonID int
}
