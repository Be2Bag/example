package ports

import "github.com/Be2Bag/example/model"

// Repository คืออินเตอร์เฟซสำหรับที่เก็บข้อมูลผู้ใช้
type RegisterRepository interface {
	CreateUser(user *model.User) error
	GetUserByEmail(email string) (*model.User, error)
	GetUsers() ([]*model.User, error)
}
