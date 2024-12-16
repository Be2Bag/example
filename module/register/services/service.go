package services

import (
	"errors"

	"github.com/Be2Bag/example/model"
	"github.com/Be2Bag/example/module/register/dto"
	"github.com/Be2Bag/example/module/register/ports"
	util "github.com/Be2Bag/example/pkg/crypto"
)

var ErrUserAlreadyExists = errors.New("ผู้ใช้มีอยู่แล้ว")

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
	if err != nil {
		return dto.RegisterResponse{}, err
	}

	// สร้างผู้ใช้ใหม่
	user := &model.User{
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
		ID:       user.ID.Hex(),
		Username: user.Username,
		Email:    user.Email,
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
			ID:        user.ID.Hex(),
			Username:  user.Username,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		})
	}

	return responses, nil
}
