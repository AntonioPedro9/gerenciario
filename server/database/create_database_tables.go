package database

import (
	"server/models"

	"gorm.io/gorm"
)

func CreateDatabaseTables(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Customer{})
	db.AutoMigrate(&models.CustomerBudget{})
	db.AutoMigrate(&models.Job{})
	db.AutoMigrate(&models.Budget{})
	db.AutoMigrate(&models.BudgetJob{})
}
