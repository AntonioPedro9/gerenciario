package database

import (
	"server/models"

	"gorm.io/gorm"
)

func ClearTestDatabase(test_db *gorm.DB) {
	test_db.Migrator().DropTable(&models.User{})
	test_db.AutoMigrate(&models.User{})

	test_db.Migrator().DropTable(&models.Client{})
	test_db.AutoMigrate(&models.Client{})
}
