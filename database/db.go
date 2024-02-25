package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/br-jeff/go-gin-simple/models"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	connectionString := "host=localhost user=dev password=dev dbname=dev port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString))

	if err != nil {
		log.Panic("erro while try to connect database, pleaase run docker compose up before")
	}

	DB.AutoMigrate(&models.Player{})

}
