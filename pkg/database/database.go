package database

import (
	"be-batch/pkg/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	config := config.GetConfig()
	psqlConn := fmt.Sprintf("host=%s user=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Bangkok", config.Database.Host, config.Database.Username, config.Database.DatabaseName, config.Database.Port)

	var err error
	db, err = gorm.Open(postgres.Open(psqlConn), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func GetDatabase() *gorm.DB {
	return db
}