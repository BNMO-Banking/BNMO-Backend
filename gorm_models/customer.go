package gormmodels

import (
	"BNMO/enum"

	"github.com/google/uuid"
)

type Customer struct {
	Base
	AccountNumber      string             `gorm:"unique"`
	Pin                string             `gorm:"unique"`
	CardNumber         string             `gorm:"unique"`
	Balance            float32            `gorm:"not_null; default:0" sql:"type:decimal(12, 2)"`
	Status             enum.AccountStatus `gorm:"not_null; default:'PENDING'"`
	PhoneNumber        string             `gorm:"not_null"`
	IdCardPath         string             `gorm:"not_null"`
	ProfilePicturePath string
	Account            Account
	AccountID          uuid.UUID
	Associates         []*Customer `gorm:"many2many:customer_associates"`
}
