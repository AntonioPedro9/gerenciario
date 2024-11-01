package database

import (
	"server/internals/models"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	_ = db.Migrator().DropTable(&models.User{}, &models.Customer{})

	err := db.AutoMigrate(&models.User{}, &models.Customer{})
	if err != nil {
		log.Fatal("Failed to run database migrations:", err)
	}
}
