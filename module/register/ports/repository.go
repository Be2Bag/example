package ports

import "github.com/Be2Bag/example/model"

type RegisterRepository interface {
	CreateUser(user *model.User) error
	GetUserByEmail(email string) (*model.User, error)
	GetUsers() ([]*model.User, error)
	GetUserByID(user_id string) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(user_id string) error
}
