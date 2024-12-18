package services

import (
	"errors"

	"github.com/Be2Bag/example/model"
	"github.com/Be2Bag/example/module/register/dto"
	registerports "github.com/Be2Bag/example/module/register/ports"
	pkgports "github.com/Be2Bag/example/pkg/ports"
	"github.com/google/uuid"
)

var ErrStaffAlreadyExists = errors.New("ผู้ใช้มีอยู่แล้ว")
var ErrStaffNotFound = errors.New("ไม่พบผู้ใช้ตาม ID ที่ระบุ")

type RegisterService struct {
	repository    registerports.RegisterRepository
	cryptoService pkgports.CryptoService
}

func NewRegisterService(repository registerports.RegisterRepository, cryptoService pkgports.CryptoService) registerports.RegisterService {
	return &RegisterService{
		repository:    repository,
		cryptoService: cryptoService,
	}
}

func (s *RegisterService) Register(req dto.RegisterRequest) (dto.RegisterResponse, error) {
	existingStaff, err := s.repository.GetStaffByEmail(req.Email)
	if err == nil && existingStaff != nil {
		return dto.RegisterResponse{}, ErrStaffAlreadyExists
	}

	hashedPassword := s.cryptoService.HasPwHelper(req.Password)

	staff := &model.Staff{
		UserID:    uuid.New().String(),
		Username:  req.Username,
		Email:     req.Email,
		Password:  hashedPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	err = s.repository.CreateStaff(staff)
	if err != nil {
		return dto.RegisterResponse{}, err
	}

	return dto.RegisterResponse{
		UserID:    staff.UserID,
		Username:  staff.Username,
		Email:     staff.Email,
		FirstName: staff.FirstName,
		LastName:  staff.LastName,
	}, nil
}

func (s *RegisterService) GetStaffs() ([]dto.RegisterResponse, error) {

	staffs, err := s.repository.GetStaffs()
	if err != nil {
		return nil, err
	}

	var responses []dto.RegisterResponse
	for _, staff := range staffs {
		responses = append(responses, dto.RegisterResponse{
			UserID:    staff.UserID,
			Username:  staff.Username,
			Email:     staff.Email,
			FirstName: staff.FirstName,
			LastName:  staff.LastName,
		})
	}

	return responses, nil
}

func (s *RegisterService) GetStaffByID(id string) (dto.RegisterResponse, error) {

	staff, err := s.repository.GetStaffByID(id)

	if staff == nil {
		return dto.RegisterResponse{}, ErrStaffNotFound
	}

	if err != nil {
		return dto.RegisterResponse{}, err
	}

	var responses = dto.RegisterResponse{
		UserID:    staff.UserID,
		Username:  staff.Username,
		Email:     staff.Email,
		FirstName: staff.FirstName,
		LastName:  staff.LastName,
	}

	return responses, nil
}

func (s *RegisterService) UpdateStaff(req dto.UpdateStaffRequest) (dto.RegisterResponse, error) {
	staff, err := s.repository.GetStaffByID(req.UserID)
	if staff == nil {
		return dto.RegisterResponse{}, ErrStaffNotFound
	}

	if err != nil {
		return dto.RegisterResponse{}, err
	}

	staff.Username = req.Username
	staff.Email = req.Email
	staff.FirstName = req.FirstName
	staff.LastName = req.LastName

	err = s.repository.UpdateStaff(staff)
	if err != nil {
		return dto.RegisterResponse{}, err
	}

	return dto.RegisterResponse{
		UserID:    staff.UserID,
		Username:  staff.Username,
		Email:     staff.Email,
		FirstName: staff.FirstName,
		LastName:  staff.LastName,
	}, nil
}

func (s *RegisterService) DeleteStaff(id string) error {
	staff, err := s.repository.GetStaffByID(id)
	if staff == nil {
		return ErrStaffNotFound
	}

	if err != nil {
		return err
	}

	err = s.repository.DeleteStaff(id)
	if err != nil {
		return err
	}

	return nil
}
