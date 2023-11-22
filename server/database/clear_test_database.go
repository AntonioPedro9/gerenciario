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

	test_db.Migrator().DropTable(&models.Service{})
	test_db.AutoMigrate(&models.Service{})

	test_db.Migrator().DropTable(&models.Budget{})
	test_db.AutoMigrate(&models.Budget{})

	test_db.Migrator().DropTable(&models.BudgetService{})
	test_db.AutoMigrate(&models.BudgetService{})

	test_db.Migrator().DropTable(&models.Appointment{})
	test_db.AutoMigrate(&models.Appointment{})
}
