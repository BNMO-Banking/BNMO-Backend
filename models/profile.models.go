package models

import "BNMO/enum"

type ProfileRes struct {
	AccountNumber      string           `json:"account_number"`
	AccountType        enum.AccountType `json:"account_type"`
	Email              string           `json:"email"`
	Username           string           `json:"username"`
	FirstName          string           `json:"first_name"`
	LastName           string           `json:"last_name"`
	CardNumber         string           `json:"card_number"`
	Balance            float32          `json:"balance"`
	PhoneNumber        string           `json:"phone_number"`
	ProfilePicturePath string           `json:"profile_pic_path"`
	AddressLine1       string           `json:"address_line_1"`
	AddressLine2       string           `json:"address_line_2"`
	City               string           `json:"city"`
	State              string           `json:"state"`
	PostalCode         string           `json:"postal_code"`
	Country            string           `json:"country"`
}
