package database

import (
	"server/models"

	"gorm.io/gorm"
)

func ClearTestDatabase(test_db *gorm.DB) {
	test_db.Migrator().DropTable(&models.User{})
	test_db.AutoMigrate(&models.User{})

	test_db.Migrator().DropTable(&models.Customer{})
	test_db.AutoMigrate(&models.Customer{})

	test_db.Migrator().DropTable(&models.Job{})
	test_db.AutoMigrate(&models.Job{})

	test_db.Migrator().DropTable(&models.Budget{})
	test_db.AutoMigrate(&models.Budget{})

	test_db.Migrator().DropTable(&models.BudgetJob{})
	test_db.AutoMigrate(&models.BudgetJob{})
}
