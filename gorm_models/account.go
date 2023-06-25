package gormmodels

import "BNMO/enum"

type Account struct {
	Base
	AccountNumber string           `json:"account_number" gorm:"unique"`
	Email         string           `json:"email" gorm:"unique; not_null"`
	Username      string           `json:"username" gorm:"unique; not_null"`
	FirstName     string           `json:"first_name" gorm:"not_null"`
	LastName      string           `json:"last_name"`
	Password      []byte           `gorm:"not_null"`
	AccountType   enum.AccountType `gorm:"not_null"`
}
