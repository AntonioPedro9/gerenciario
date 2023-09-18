package database

import "server/models"

func CreateDatabaseTables() {
	DB.AutoMigrate(&models.User{})
}
