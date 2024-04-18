package entities

type PokemonItem struct {
	ID        int `gorm:"primaryKey"`
	ItemID    int
	Name      string
	Cost      int
	SpriteURL string
}
