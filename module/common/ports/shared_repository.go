package ports

import "github.com/Be2Bag/example/model"

type SharedRepository interface {
	GetStaffByEmail(email string) (*model.Staff, error)
	GetStaffByID(staff_id string) (*model.Staff, error)
}
