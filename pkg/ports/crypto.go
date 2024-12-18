package ports

type CryptoService interface {
	GenerateJWTToken(data map[string]interface{}) (string, error)
	ValidateJWTToken(tokenStr string) (map[string]interface{}, error)
	HasPwHelper(pw string) string
	Encrypt(data string) (string, error)
	Decrypt(data string) (string, error)
}
