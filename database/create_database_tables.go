package database

import (
	"server/models"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func CreateDatabaseTables(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Client{})

	log.Info("Database tables created")
}
