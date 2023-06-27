package models

import (
	"BNMO/enum"
	"mime/multipart"

	"github.com/google/uuid"
)

type RegisterReq struct {
	FirstName       string                `form:"first_name"`
	LastName        string                `form:"last_name"`
	Email           string                `form:"email"`
	Username        string                `form:"username"`
	PhoneNumber     string                `form:"phone_number"`
	AddressLine1    string                `form:"address_line1"`
	AddressLine2    string                `form:"address_line2"`
	City            string                `form:"city"`
	State           string                `form:"state"`
	PostalCode      string                `form:"postal_code"`
	Country         string                `form:"country"`
	IdCard          *multipart.FileHeader `form:"id_card"`
	Password        string                `form:"password"`
	ConfirmPassword string                `form:"confirm_password"`
}

type LoginReq struct {
	EmailUsername string `json:"email_username"`
	Password      string `json:"password"`
}

type LoginResAccount struct {
	Id          uuid.UUID        `json:"id"`
	Email       string           `json:"email"`
	Username    string           `json:"username"`
	AccountType enum.AccountType `json:"account_type"`
	AccountRole enum.AccountRole `json:"account_role"`
}

type LoginRes struct {
	Account   LoginResAccount `json:"account"`
	PinStatus enum.PinStatus  `json:"pin_status"`
	Token     string          `json:"token"`
}
