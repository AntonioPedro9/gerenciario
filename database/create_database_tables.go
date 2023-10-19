package database

import (
	"server/models"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func CreateDatabaseTables(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Client{})
	db.AutoMigrate(&models.Service{})
	db.AutoMigrate(&models.Budget{})
	db.AutoMigrate(&models.BudgetService{})

	log.Info("Database tables created")
}
