package services

import (
	"errors"

	"github.com/Be2Bag/example/model"
	"github.com/Be2Bag/example/module/register/dto"
	"github.com/Be2Bag/example/module/register/ports"
	util "github.com/Be2Bag/example/pkg/crypto"
	"github.com/google/uuid"
)

var ErrUserAlreadyExists = errors.New("ผู้ใช้มีอยู่แล้ว")
var ErrUserNotFound = errors.New("ไม่พบผู้ใช้ตาม ID ที่ระบุ")

type RegisterService struct {
	repository ports.RegisterRepository
}

func NewRegisterService(repository ports.RegisterRepository) ports.RegisterService {
	return &RegisterService{
		repository: repository,
	}
}

func (service *RegisterService) Register(req dto.RegisterRequest) (dto.RegisterResponse, error) {
	existingUser, err := service.repository.GetUserByEmail(req.Email)
	if err == nil && existingUser != nil {
		return dto.RegisterResponse{}, ErrUserAlreadyExists
	}

	hashedPassword := util.HasPwHelper(req.Password)

	user := &model.User{
		UserID:    uuid.New().String(),
		Username:  req.Username,
		Email:     req.Email,
		Password:  hashedPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	err = service.repository.CreateUser(user)
	if err != nil {
		return dto.RegisterResponse{}, err
	}

	return dto.RegisterResponse{
		UserID:    user.UserID,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, nil
}

func (service *RegisterService) GetUsers() ([]dto.RegisterResponse, error) {

	users, err := service.repository.GetUsers()
	if err != nil {
		return nil, err
	}

	var responses []dto.RegisterResponse
	for _, user := range users {
		responses = append(responses, dto.RegisterResponse{
			UserID:    user.UserID,
			Username:  user.Username,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		})
	}

	return responses, nil
}

func (service *RegisterService) GetUserByID(id string) (dto.RegisterResponse, error) {

	user, err := service.repository.GetUserByID(id)

	if user == nil {
		return dto.RegisterResponse{}, ErrUserNotFound
	}

	if err != nil {
		return dto.RegisterResponse{}, err
	}

	var responses = dto.RegisterResponse{
		UserID:    user.UserID,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	return responses, nil
}

func (service *RegisterService) UpdateUser(req dto.UpdateUserRequest) (dto.RegisterResponse, error) {
	user, err := service.repository.GetUserByID(req.UserID)
	if user == nil {
		return dto.RegisterResponse{}, ErrUserNotFound
	}

	if err != nil {
		return dto.RegisterResponse{}, err
	}

	user.Username = req.Username
	user.Email = req.Email
	user.FirstName = req.FirstName
	user.LastName = req.LastName

	err = service.repository.UpdateUser(user)
	if err != nil {
		return dto.RegisterResponse{}, err
	}

	return dto.RegisterResponse{
		UserID:    user.UserID,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, nil
}

func (service *RegisterService) DeleteUser(id string) error {
	user, err := service.repository.GetUserByID(id)
	if user == nil {
		return ErrUserNotFound
	}

	if err != nil {
		return err
	}

	err = service.repository.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}
