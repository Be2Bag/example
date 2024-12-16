package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func Encrypt(data string) (string, error) {

	key := os.Getenv("KEY_SECRET")

	if len(key) != 32 {
		return "", errors.New("invalid key size")
	}
	iv := make([]byte, aes.BlockSize)
	for i := range iv {
		iv[i] = 0
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	padding := aes.BlockSize - len(jsonData)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	paddedData := append(jsonData, padtext...)

	encrypted := make([]byte, len(paddedData))
	mode.CryptBlocks(encrypted, paddedData)

	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func Decrypt(data string) (string, error) {

	key := os.Getenv("KEY_SECRET")

	if len(key) != 32 {
		return "", errors.New("invalid key size")
	}
	iv := make([]byte, aes.BlockSize)
	for i := range iv {
		iv[i] = 0
	}

	encryptedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(encryptedData))
	mode.CryptBlocks(decrypted, encryptedData)

	padding := decrypted[len(decrypted)-1]
	if int(padding) > aes.BlockSize || int(padding) == 0 {
		return "", errors.New("invalid padding size")
	}
	decrypted = decrypted[:len(decrypted)-int(padding)]

	var result string
	err = json.Unmarshal(decrypted, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}

func HasPwHelper(pw string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pw), 10)
	if err != nil {
		return err.Error()
	}
	return string(hashedPassword)
}
