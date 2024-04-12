package entities

type Pokemon struct {
	ID                                   int `gorm:"primaryKey"`
	Name                                 string
	SpriteFrontDefaultShowdownURL        string
	SpriteFrontDefaultOfficialArtworkURL string
	Height                               float64
	Weight                               float64
	BaseExperience                       int
	MinimumLevel                         int
	EvolvedPokemonID                     *int
}
