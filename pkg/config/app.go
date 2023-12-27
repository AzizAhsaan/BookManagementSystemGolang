package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db * gorm.DB
)
func Connect() {
	d, err := gorm.Open("postgres", "user=user dbname=dbname sslmode=disable password=password")

	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}