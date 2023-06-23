package gormmodels

import "github.com/google/uuid"

type CustomerAddress struct {
	Base
	AddressLine1 string   `json:"address_line_1" gorm:"not_null"`
	AddressLine2 string   `json:"address_line_2"`
	City         string   `json:"city" gorm:"not_null"`
	State        string   `json:"state" gorm:"not_null"`
	PostalCode   string   `json:"postal_code" gorm:"not_null"`
	Country      string   `json:"country" gorm:"not_null"`
	Customer     Customer `json:"customer"`
	customerID   uuid.UUID
}