package db

import (
	"github.com/timileyin19/gofarm/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var env = config.EnVar

func ConnectToDb() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(env.GetString("DNS")), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db, err
}