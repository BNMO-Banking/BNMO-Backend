package gormmodels

import "github.com/google/uuid"

type Request struct {
	Base
	RequestType     string   `json:"request_type" gorm:"not_null"`
	Currency        string   `json:"currency" gorm:"not_null"`
	Amount          int64    `json:"amount" gorm:"not_null"`
	ConvertedAmount float32  `json:"converted_amount" gorm:"not_null" sql:"type:decimal(12, 2)"`
	Status          string   `json:"request_status" gorm:"not_null; default:'PENDING"`
	Remarks         string   `json:"remarks" sql:"type:text"`
	Customer        Customer `json:"customer"`
	customerID      uuid.UUID
}
