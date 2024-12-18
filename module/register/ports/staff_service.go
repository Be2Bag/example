package ports

import (
	"github.com/Be2Bag/example/module/register/dto"
)

type RegisterService interface {
	Register(req dto.RegisterRequest) (dto.RegisterResponse, error)
	GetStaffs() ([]dto.RegisterResponse, error)
	GetStaffByID(staff_id string) (dto.RegisterResponse, error)
	UpdateStaff(req dto.UpdateStaffRequest) (dto.RegisterResponse, error)
	DeleteStaff(id string) error
}
