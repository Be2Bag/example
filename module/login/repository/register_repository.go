// module/register/repository/register_repository.go
package repository

import (
	"github.com/Be2Bag/example/model"
	"github.com/Be2Bag/example/module/register/ports"
	"gorm.io/gorm"
)

type registerRepository struct {
    db *gorm.DB
}

func NewRegisterRepository(db *gorm.DB) ports.RegisterRepository {
    return &registerRepository{
        db: db,
    }
}

func (r *registerRepository) CreateUser(user *model.User) error {
    return r.db.Create(user).Error
}

func (r *registerRepository) GetUserByEmail(email string) (*model.User, error) {
    var user model.User
    result := r.db.Where("email = ?", email).First(&user)
    if result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}
