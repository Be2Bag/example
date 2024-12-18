package util

import (
	"regexp"

	"github.com/Be2Bag/example/pkg/ports"
	"github.com/go-playground/validator/v10"
)

type validatorService struct {
	validate *validator.Validate
}

func NewValidatorService() ports.ValidatorService {
	v := validator.New()
	v.RegisterValidation("datavalid", customPasswordValidator)
	return &validatorService{validate: v}
}

func (v *validatorService) ValidatePassword(password string) bool {
	err := v.validate.Var(password, "datavalid")
	return err == nil
}

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
