package models

type ContractService struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	ContractID uint
	ServiceID  uint
}
