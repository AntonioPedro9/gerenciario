package database

import (
	"server/models"

	"gorm.io/gorm"
)

func CreateDatabaseTables(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Client{})
	db.AutoMigrate(&models.Service{})
	db.AutoMigrate(&models.Budget{})
	db.AutoMigrate(&models.BudgetService{})
	db.AutoMigrate(&models.Appointment{})
}
