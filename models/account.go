package models

import (
	"database/sql"
	"mime/multipart"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// GORM models
type Account struct {
	gorm.Model
	IsAdmin			sql.NullBool 	`json:"is_admin" gorm:"not null;default:false"`
	AccountStatus	sql.NullString	`json:"account_status" gorm:"not null;default:'pending'"`
	FirstName		string	`json:"first_name" gorm:"not null"`
	LastName		string	`json:"last_name" gorm:"not null"`
    Email 			string	`json:"email" gorm:"unique; not null"`
    Username 		string	`json:"username" gorm:"unique; not null"`
    Password 		[]byte	`json:"-" gorm:"not null"`
    ImagePath 		string	`json:"image_path" gorm:"not null"`
	AccountNumber	string	`json:"account_number" gorm:"unique"`
	Pin				[]byte	`json:"-"`
    Balance 		int64	`json:"balance" gorm:"not null;default:0"`
	Transfers		[]Transfer		`json:"-" gorm:"foreignKey:SourceID"`
	Requests		[]Request		`json:"-" gorm:"foreignKey:DestinationID"`
	TransferDestination []*Account `json:"-" gorm:"many2many:transfer_destination"`
}

// JSON models
type RegisterRequest struct {
	FirstName	string	`form:"first_name"`
	LastName	string	`form:"last_name"`
	Username	string	`form:"username"`
	Email		string	`form:"email"`
	Password	string	`form:"password"`
	Image		*multipart.FileHeader	`form:"image"`
}

type LoginRequest struct {
	Email		string 	`json:"email"`
	Password	string	`json:"password"`
}

type ValidateAccount struct {
	Id			uint	`json:"id"`
	Validation	string	`json:"validation"`
}

type ChangeImageRequest struct {
	Id			string	`form:"id"`
	OldName		string	`form:"old_name"`
	NewImage	*multipart.FileHeader	`form:"new_image"`
}
type ChangePinRequest struct {
	Id			uint	`json:"id"`
	Pin			string	`json:"pin"`
	RepeatPin	string	`json:"repeat_pin"`
}

type ChangePassRequest struct {
	Id			uint	`json:"id"`
	OldPass		string	`json:"old_pass"`
	NewPass		string	`json:"new_pass"`
	RepeatPass	string	`json:"repeat_pass"`
}

type EditDataRequest struct {
	Id			uint	`json:"id"`
	FirstName	string	`json:"first_name"`
	LastName	string	`json:"last_name"`
	Username	string	`json:"username"`
	Email		string	`json:"email"`
	Balance 	int64	`json:"balance"`
}

type ResetPinRequest struct {
	Id			uint	`json:"id"`
}

type DeleteAccRequest struct {
	Id			uint	`json:"id"`
}

// Function to hash password using bcrypt with salt
func (account *Account) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	account.Password = hashedPassword
}

// Function to hash pin using bcrypt with salt
func (account *Account) SetPin(pin string) {
	hashedPin, _ := bcrypt.GenerateFromPassword([]byte(pin), 8)
	account.Pin = hashedPin
}

// Function to compare user inputted password with the one inside the database
func (account *Account) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(account.Password, []byte(password))
}

// Function to compare user inputted pin with the one inside the database
func (account *Account) ComparePin(pin string) error {
	return bcrypt.CompareHashAndPassword(account.Pin, []byte(pin))
}