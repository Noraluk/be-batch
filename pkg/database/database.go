package database

import (
	pokemonEtt "be-batch/jobs/pokemon/entities"
	pokemonItemEtt "be-batch/jobs/pokemon_item/entities"
	"be-batch/pkg/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	config := config.GetConfig()
	psqlConn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Bangkok", config.Database.Host, config.Database.Username, config.Database.Password, config.Database.DatabaseName, config.Database.Port)

	var err error
	db, err = gorm.Open(postgres.Open(psqlConn), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&pokemonEtt.Pokemon{})
	db.AutoMigrate(&pokemonEtt.PokemonType{})
	db.AutoMigrate(&pokemonEtt.PokemonAbility{})
	db.AutoMigrate(&pokemonEtt.PokemonWeakness{})
	db.AutoMigrate(&pokemonEtt.PokemonStat{})
	db.AutoMigrate(&pokemonItemEtt.PokemonItem{})

	return nil
}

func GetDatabase() *gorm.DB {
	return db
}
