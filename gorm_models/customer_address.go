package gormmodels

import "github.com/google/uuid"

type CustomerAddress struct {
	Base
	AddressLine1 string `gorm:"not_null"`
	AddressLine2 string
	City         string `gorm:"not_null"`
	State        string `gorm:"not_null"`
	PostalCode   string `gorm:"not_null"`
	Country      string `gorm:"not_null"`
	Customer     Customer
	customerID   uuid.UUID
}
