package helpers

import (
	"errors"
	"strings"
)

func IsValidEmail(email string) error {
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return errors.New("Invalid email format")
	}
	return nil
}

func ValidateBalance(balance int) error {
	if balance < 0 || balance > 100000000 {
		return errors.New("Balance must be between 0 and 100,000,000")
	}
	return nil
}

func ValidatePrice(price int) error {
	if price < 0 || price > 50000000 {
		return errors.New("Price must be between 0 and 50,000,000")
	}
	return nil
}
