package services

import (
	"errors"

	"github.com/Be2Bag/example/module/session/dto"
	"github.com/Be2Bag/example/module/session/ports"
	util "github.com/Be2Bag/example/pkg/crypto"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidEmail = errors.New("อีเมล์ไม่ถูกต้อง")
var ErrInvalidPassword = errors.New("รหัสผ่านไม่ถูกต้อง")

type SessionService struct {
	repository ports.SessionRepository
}

func NewSessionService(repository ports.SessionRepository) ports.SessionService {
	return &SessionService{
		repository: repository,
	}
}

func (s *SessionService) Login(sessionRequest dto.SessionRequest) (string, error) {

	user, err := s.repository.GetUserByEmail(sessionRequest.Email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", ErrInvalidEmail
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(sessionRequest.Password)); err != nil {
		return "", ErrInvalidPassword
	}

	data := map[string]interface{}{
		"userID": user.UserID,
		"email":  user.Email,
	}

	token, errToken := util.GenerateJWTToken(data)

	if errToken != nil {
		return "", errToken
	}

	return token, nil
}
