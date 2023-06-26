package utils

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func ValidateEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9. %+\-]+@[a-z0-9. %+\-]+\.[a-z0-9. %+\-]`)
	return Re.MatchString(email)
}

func ValidatePhoneNumber(phone string) bool {
	Re := regexp.MustCompile(`^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,6}$`)
	return Re.MatchString(phone)
}

// Expected account number XXX-XXX-XXXX
func GenerateAccountNumber() string {
	const format = "%03d%03d%04d"
	firstSequel := rand.Intn(1000)
	secondSequel := rand.Intn(1000)
	thirdSequel := rand.Intn(10000)

	return fmt.Sprintf(format, firstSequel, secondSequel, thirdSequel)
}

// Expected card number XXXX-XXXX-XXXX-XXXX
func GenerateCardNumber() string {
	const format = "%04d%04d%04d%04d"
	firstSequel := rand.Intn(10000)
	secondSequel := rand.Intn(10000)
	thirdSequel := rand.Intn(10000)
	fourthSequel := rand.Intn(10000)

	return fmt.Sprintf(format, firstSequel, secondSequel, thirdSequel, fourthSequel)
}

func HashPassword(password string) ([]byte, error) {
	salt, err := strconv.Atoi(os.Getenv("PASS_SALT"))
	if err != nil {
		return nil, err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	if err != nil {
		return nil, err
	}

	return hashed, nil
}

func ComparePassword(password1 []byte, password2 string) error {
	return bcrypt.CompareHashAndPassword(password1, []byte(password2))
}
