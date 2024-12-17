package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"os"

	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(data string) (string, error) {

	key := os.Getenv("KEY_SECRET")

	if len(key) != 32 {
		return "", errors.New("ขนาดกุญแจไม่ถูกต้อง")
	}
	iv := make([]byte, aes.BlockSize)
	_, err := rand.Read(iv)
	if err != nil {
		return "", errors.New("ไม่สามารถสร้าง IV ได้")
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
		return "", errors.New("ขนาดกุญแจไม่ถูกต้อง")
	}
	iv := make([]byte, aes.BlockSize)
	_, err := rand.Read(iv)
	if err != nil {
		return "", errors.New("ไม่สามารถสร้าง IV ได้")
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
		return "", errors.New("ขนาดการเติมไม่ถูกต้อง")
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
		return "เกิดข้อผิดพลาดในการสร้างรหัสผ่าน"
	}
	return string(hashedPassword)
}

func GenerateJWTToken(data map[string]interface{}) (string, error) {

	claims := jwt.MapClaims{
		"data": data,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("ไม่พบคีย์ลับ JWT")
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWTToken(tokenStr string) (map[string]interface{}, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, errors.New("ไม่พบคีย์ลับ JWT")
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		data, ok := claims["data"].(map[string]interface{})
		if !ok {
			return nil, errors.New("ไม่พบข้อมูลในโทเค็น")
		}
		return data, nil
	}

	return nil, errors.New("โทเค็นไม่ถูกต้อง")
}
