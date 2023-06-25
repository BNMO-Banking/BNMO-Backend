package gormmodels

import "BNMO/enum"

type Account struct {
	Base
	Email       string `gorm:"unique; not_null"`
	Username    string `gorm:"unique; not_null"`
	FirstName   string `gorm:"not_null"`
	LastName    string
	Password    []byte           `gorm:"not_null"`
	AccountType enum.AccountType `gorm:"not_null"`
}
