package models

type BudgetOffering struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	ContractID uint
	ServiceID  uint
}
