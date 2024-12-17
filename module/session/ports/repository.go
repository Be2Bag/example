package ports

import "github.com/Be2Bag/example/model"

type SessionRepository interface {
	GetUserByEmail(email string) (*model.User, error)
}
