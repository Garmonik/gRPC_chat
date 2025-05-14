package validate_lib

import (
	passwordvalidator "github.com/wagslane/go-password-validator"
	"net/mail"
	"strconv"
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

func ConversionStringToUint(str string) (uint, error) {
	num, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	result := uint(num)
	return result, nil
}
