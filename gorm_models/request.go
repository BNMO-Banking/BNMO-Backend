package gormmodels

import "github.com/google/uuid"

type Request struct {
	Base
	RequestType     string  `gorm:"not_null"`
	Currency        string  `gorm:"not_null"`
	Amount          int64   `gorm:"not_null"`
	ConvertedAmount float32 `gorm:"not_null" sql:"type:decimal(12, 2)"`
	Status          string  `gorm:"not_null; default:'PENDING"`
	Remarks         string  `sql:"type:text"`
	Customer        Customer
	customerID      uuid.UUID
}
