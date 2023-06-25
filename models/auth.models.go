package models

import (
	"BNMO/enum"
	"mime/multipart"
)

type RegisterReq struct {
	FirstName       string                `json:"first_name"`
	LastName        string                `json:"last_name"`
	AddressLine1    string                `json:"address_line1"`
	AddressLine2    string                `json:"address_line2"`
	City            string                `json:"city"`
	State           string                `json:"state"`
	PostalCode      string                `json:"postal_code"`
	Country         string                `json:"country"`
	Email           string                `json:"email"`
	Username        string                `json:"username"`
	PhoneNumber     string                `json:"phone_number"`
	IdCard          *multipart.FileHeader `json:"id_card"`
	Password        string                `json:"password"`
	ConfirmPassword string                `json:"confirm_password"`
}

type LoginReq struct {
	EmailUsername string `json:"email_username"`
	Password      string `json:"password"`
}

type LoginResAccount struct {
	Email       string           `json:"email"`
	Username    string           `json:"username"`
	AccountType enum.AccountType `json:"account_type"`
	AccountRole enum.AccountRole `json:"account_role"`
}

type LoginRes struct {
	Account LoginResAccount `json:"account"`
	Token   string          `json:"token"`
}
