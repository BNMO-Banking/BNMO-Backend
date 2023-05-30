package models

import (
	"gorm.io/gorm"
)

type Request struct {
	gorm.Model
	DestinationID    uint		`json:"destination_id"`
	RequestType		string		`json:"request_type" gorm:"not null"`
	Amount         	int64		`json:"amount" gorm:"not null"`
	Currency		string		`json:"currency" gorm:"not null"`
	ConvertedAmount	int64		`json:"converted_amount"`
	Status			string		`json:"status" gorm:"not null;default:'pending'"`
	Destination		Account		`json:"destination" gorm:"foreignKey:DestinationID;references:ID"`
}

type ValidateRequest struct {
	RequestID		uint		`json:"request_id"`
	Status			string		`json:"status"`
}