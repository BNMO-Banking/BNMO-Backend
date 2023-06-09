package utils

import (
	"fmt"
	"os"
	"strconv"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashPin(id uuid.UUID, pin string) ([]byte, error) {
	combined := CombinePin(id, pin)

	salt, err := strconv.Atoi(os.Getenv("PIN_SALT"))
	if err != nil {
		return nil, err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(combined), salt)
	if err != nil {
		return nil, err
	}

	return hashed, nil
}

func ComparePin(pin1 []byte, pin2 string) error {
	return bcrypt.CompareHashAndPassword(pin1, []byte(pin2))
}

func CombinePin(id uuid.UUID, pin string) string {
	return fmt.Sprintf(os.Getenv("PIN_FORMAT"), id.String(), pin)
}
