package gormmodels

import (
	"BNMO/enum"

	"github.com/google/uuid"
)

type Customer struct {
	Base
	Pin                string             `gorm:"unique"`
	CardNumber         string             `json:"card_number" gorm:"unique"`
	Balance            float32            `json:"balance" gorm:"not_null; default:0" sql:"type:decimal(12, 2)"`
	Status             enum.AccountStatus `json:"account_status" gorm:"not_null; default:'PENDING'"`
	PhoneNumber        string             `json:"phone_number" gorm:"not_null"`
	IdCardPath         string             `json:"id_card_path" gorm:"not_null"`
	ProfilePicturePath string             `json:"profile_pic_path"`
	Account            Account            `json:"account"`
	AccountID          uuid.UUID
	Associates         []*Customer `gorm:"many2many:customer_associates"`
}
