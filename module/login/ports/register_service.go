// module/register/ports/register_service.go
package ports

import "github.com/Be2Bag/example/module/register/dto"

type RegisterService interface {
    Register(req dto.RegisterRequest) (dto.RegisterResponse, error)
}
