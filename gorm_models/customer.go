package gormmodels

import (
	"BNMO/enum"

	"github.com/google/uuid"
)

type Customer struct {
	Base
	AccountNumber      string             `gorm:"unique; default:null"`
	Pin                string             `gorm:"unique; default:null"`
	CardNumber         string             `gorm:"unique; default:null"`
	Balance            float32            `gorm:"not null; default:0" sql:"type:decimal(12, 2)"`
	Status             enum.AccountStatus `gorm:"not null; default:'PENDING'"`
	PhoneNumber        string             `gorm:"not null"`
	IdCardPath         string             `gorm:"not null"`
	ProfilePicturePath string             `gorm:"default:null"`
	Account            Account
	AccountID          uuid.UUID
	Address            CustomerAddress `gorm:"foreignKey:AddressID"`
	AddressID          uuid.UUID
	Associates         []*Customer `gorm:"many2many:customer_associates"`
}
