package models

import (
	"github.com/google/uuid"
)

type AddAssociatesReq struct {
	Id                uuid.UUID `json:"id"`
	DestinationNumber string    `json:"destination_number"`
}

type DestinationsRes struct {
	AccountNumber string `json:"account_number"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
}

type DestinationResList struct {
	Data []DestinationsRes `json:"data"`
}
