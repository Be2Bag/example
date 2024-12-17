package ports

import "github.com/Be2Bag/example/model"

type SharedRepository interface {
	GetUserByEmail(email string) (*model.User, error)
	GetUserByID(user_id string) (*model.User, error)
}
