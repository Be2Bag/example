// module/register/ports/register_repository.go
package ports

import "github.com/Be2Bag/example/model"

type RegisterRepository interface {
    CreateUser(user *model.User) error
    GetUserByEmail(email string) (*model.User, error)
}
