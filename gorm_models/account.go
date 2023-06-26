package gormmodels

import "BNMO/enum"

type Account struct {
	Base
	Email       string           `gorm:"unique; not null"`
	Username    string           `gorm:"unique; not null"`
	FirstName   string           `gorm:"not null"`
	LastName    string           `gorm:"default:null"`
	Password    []byte           `gorm:"not null"`
	AccountType enum.AccountType `gorm:"not null"`
}
