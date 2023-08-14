package config

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config ConfigInf) *gorm.DB {
	var dsn string = config.Get("DSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	sql, err := db.DB()

	sql.SetConnMaxIdleTime(time.Minute * 5)
	sql.SetConnMaxLifetime(time.Hour)
	sql.SetMaxIdleConns(3)
	sql.SetMaxOpenConns(10)

	return db

}
