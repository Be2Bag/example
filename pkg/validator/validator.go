package util

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func customPasswordValidator(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	var (
		uppercase = regexp.MustCompile(`[A-Z]`).MatchString
		lowercase = regexp.MustCompile(`[a-z]`).MatchString
		number    = regexp.MustCompile(`[0-9]`).MatchString
		special   = regexp.MustCompile(`[!@#~$%^&*()+|_]{1}`).MatchString
	)
	return uppercase(password) && lowercase(password) && number(password) && special(password)
}

func RegisterValidators(v *validator.Validate) {
	v.RegisterValidation("datavalid", customPasswordValidator)
}
