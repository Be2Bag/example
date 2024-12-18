package ports

type ValidatorService interface {
	ValidatePassword(password string) bool
}
