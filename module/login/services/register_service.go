// module/register/services/register_service.go
package services

import (
	"errors"

	"github.com/Be2Bag/example/model"
	"github.com/Be2Bag/example/module/register/dto"
	"github.com/Be2Bag/example/module/register/ports"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type registerService struct {
    repo ports.RegisterRepository
}

func NewRegisterService(repo ports.RegisterRepository) ports.RegisterService {
    return &registerService{
        repo: repo,
    }
}

func (s *registerService) Register(req dto.RegisterRequest) (dto.RegisterResponse, error) {
    // Check if user already exists
    existingUser, err := s.repo.GetUserByEmail(req.Email)
    if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
        return dto.RegisterResponse{}, err
    }
    if existingUser != nil {
        return dto.RegisterResponse{}, errors.New("user already exists")
    }

    // Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        return dto.RegisterResponse{}, err
    }

    // Create user
    user := model.User{
        Username: req.Username,
        Email:    req.Email,
        Password: string(hashedPassword),
    }

    err = s.repo.CreateUser(&user)
    if err != nil {
        return dto.RegisterResponse{}, err
    }

    // Prepare response
    resp := dto.RegisterResponse{
        ID:       user.ID,
        Username: user.Username,
        Email:    user.Email,
    }

    return resp, nil
}
