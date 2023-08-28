package initializers

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var TestDB *gorm.DB

func ConnectToTestDatabase() {
	connectionString := "user=postgres password=0000 host=localhost port=5432 dbname=postgres sslmode=disable"

	var err error

	TestDB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}
