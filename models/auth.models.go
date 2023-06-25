package models

import "mime/multipart"

type RegisterReq struct {
	FirstName       string                `json:"first_name"`
	LastName        string                `json:"last_name"`
	Email           string                `json:"email"`
	Username        string                `json:"username"`
	PhoneNumber     string                `json:"phone_number"`
	IdCard          *multipart.FileHeader `json:"id_card"`
	Password        string                `json:"password"`
	ConfirmPassword string                `json:"confirm_password"`
}
