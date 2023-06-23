package gormmodels

import "github.com/google/uuid"

type Admin struct {
	Base
	Role      string  `json:"role"`
	Account   Account `json:"account"`
	AccountID uuid.UUID
}