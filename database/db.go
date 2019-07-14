package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"go-nfe-repeater/configuration"
	"go-nfe-repeater/nfe"
	"log"
)

func NewDatabaseConnection() *gorm.DB {
	db, err := gorm.Open("sqlite3", configuration.GetConfiguration().Database["storage"])

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&nfe.Nfe{})

	return db
}

func NewEmbeddedDatabaseConnection() *gorm.DB {
	db, err := gorm.Open("sqlite3", ":memory:")

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&nfe.Nfe{})

	return db
}