package ports

import "github.com/Be2Bag/example/module/register/dto"

// Service คืออินเตอร์เฟซสำหรับบริการลงทะเบียน
type RegisterService interface {
	Register(req dto.RegisterRequest) (dto.RegisterResponse, error)
	GetUsers() ([]dto.RegisterResponse, error)
}
