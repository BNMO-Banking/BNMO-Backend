package models

import "github.com/google/uuid"

type PinReq struct {
	Id  uuid.UUID `json:"id"`
	Pin string    `json:"pin"`
}
