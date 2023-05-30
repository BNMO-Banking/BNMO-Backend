package models

import (
	"gorm.io/gorm"
)

type Transfer struct {
	gorm.Model
	SourceID 		uint		`json:"source_id"`
    Destination 	string		`json:"destination"`
    Amount 			int64		`json:"amount" gorm:"not null"`
	Currency		string		`json:"currency" gorm:"not null"`
	ConvertedAmount	int64		`json:"converted_amount"`
	Status			string		`json:"status" gorm:"not null;default:'pending'"`
	Source			Account		`json:"source" gorm:"foreignKey:SourceID;references:ID"`
}

type CheckDestinationRequest struct {
	DestinationNumber	string	`json:"destination_number"`
}
type AddDestinationRequest struct {
	Id			uint	`json:"user_id"`
	DestinationNumber	string	`json:"destination_number"`
}

type DestinationResponse struct {
	AccountNumber		string		`json:"account_number"`
	FirstName			string		`json:"first_name"`
	LastName			string		`json:"last_name"`
	Username			string		`json:"username"`
}