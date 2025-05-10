package validate_lib

import (
	passwordvalidator "github.com/wagslane/go-password-validator"
	"net/mail"
)

func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func ValidPassword(password string) bool {
	minEntropy := 50.0
	err := passwordvalidator.Validate(password, minEntropy)
	if err != nil {
		return false
	}
	return true
}
