package gormmodels

import "github.com/google/uuid"

type Customer struct {
	Base
	pin                string  `gorm:"unique; not_null"`
	CardNumber         string  `json:"card_number" gorm:"unique; not_null"`
	Balance            float32   `json:"balance" gorm:"not_null; default:0" sql:"type:decimal(12, 2)"`
	Status			   string  `json:"account_status" gorm:"not_null; default:'PENDING'"`
	PhoneNumber        string  `json:"phone_number" gorm:"not_null"`
	IdCardPath         string  `json:"id_card_path" gorm:"not_null"`
	ProfilePicturePath string  `json:"profile_pic_path"`
	Account            Account `json:"account"`
	accountID          uuid.UUID
	associates			[]*Customer	`gorm:"many2many:customer_associates"`
}