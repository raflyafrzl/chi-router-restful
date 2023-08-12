package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config Config) *gorm.DB {
	var dsn string = config.Get("DSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db

}
