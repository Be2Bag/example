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

// RegisterService คือบริการสำหรับการลงทะเบียนผู้ใช้ใหม่
type RegisterService struct {
	repository ports.RegisterRepository
}

// NewRegisterService สร้างบริการลงทะเบียนใหม่
func NewRegisterService(repository ports.RegisterRepository) ports.RegisterService {
	return &RegisterService{
		repository: repository,
	}
}

// Register จัดการการลงทะเบียนผู้ใช้ใหม่
func (service *RegisterService) Register(req dto.RegisterRequest) (dto.RegisterResponse, error) {
	// ตรวจสอบว่ามีผู้ใช้ที่มีอีเมลนี้อยู่แล้วหรือไม่
	existingUser, err := service.repository.GetUserByEmail(req.Email)
	if err == nil && existingUser != nil {
		return dto.RegisterResponse{}, ErrUserAlreadyExists
	}

	// แฮชพาสเวิร์ดของผู้ใช้
	hashedPassword := util.HasPwHelper(req.Password)

	// สร้างผู้ใช้ใหม่
	user := &model.User{
		UserID:    uuid.New().String(),
		Username:  req.Username,
		Email:     req.Email,
		Password:  hashedPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	// บันทึกผู้ใช้ลงในฐานข้อมูล
	err = service.repository.CreateUser(user)
	if err != nil {
		return dto.RegisterResponse{}, err
	}

	// ส่งคืนการตอบสนองหลังการลงทะเบียนสำเร็จ
	return dto.RegisterResponse{
		UserID:    user.UserID,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, nil
}

// GetUsers ดึงข้อมูลผู้ใช้ทั้งหมด
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
