package database

import (
	"os"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB, TEST_DB *gorm.DB

func ConnectToDatabase() {
	var err error

	DB, err = gorm.Open(postgres.Open(os.Getenv("CONNECTION_STRING")), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}

func ConnectToTestDatabase() {
	var err error

	TEST_DB, err = gorm.Open(postgres.Open(os.Getenv("TEST_CONNECTION_STRING")), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to test database")
	}
}
