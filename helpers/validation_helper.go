package helpers

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/mochachmar/pbi-btpns-fd-mochachmar/app"
)

var (
	ErrInvalidEmail     = errors.New("invalid email format")
	ErrPasswordTooShort = errors.New("password should be at least 6 characters long")
)

func ValidateUser(user app.User) error {
	if _, err := govalidator.ValidateStruct(user); err != nil {
		return err
	}

	if !govalidator.IsEmail(user.Email) {
		return ErrInvalidEmail
	}

	if len(user.Password) < 6 {
		return ErrPasswordTooShort
	}

	return nil
}
