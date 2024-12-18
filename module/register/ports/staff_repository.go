package ports

import "github.com/Be2Bag/example/model"

type RegisterRepository interface {
	CreateStaff(staff *model.Staff) error
	GetStaffByEmail(email string) (*model.Staff, error)
	GetStaffs() ([]*model.Staff, error)
	GetStaffByID(staff_id string) (*model.Staff, error)
	UpdateStaff(staff *model.Staff) error
	DeleteStaff(staff_id string) error
}
