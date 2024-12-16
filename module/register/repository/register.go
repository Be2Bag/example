package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Be2Bag/example/model"
	"github.com/Be2Bag/example/module/register/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// RegisterRepository คือที่เก็บข้อมูลสำหรับการลงทะเบียนผู้ใช้
type RegisterRepository struct {
	collection *mongo.Collection
}

// NewRegisterRepository สร้างที่เก็บข้อมูลใหม่
func NewRegisterRepository(db *mongo.Database) ports.RegisterRepository {
	collection := db.Collection("users")
	return &RegisterRepository{
		collection: collection,
	}
}

// CreateUser สร้างผู้ใช้ใหม่ในฐานข้อมูล
func (repository *RegisterRepository) CreateUser(user *model.User) error {
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	_, err := repository.collection.InsertOne(context.TODO(), user)
	return err
}

// GetUserByEmail ดึงข้อมูลผู้ใช้ตามอีเมล
func (repository *RegisterRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := repository.collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (repository *RegisterRepository) GetUsers() ([]*model.User, error) {
	cursor, err := repository.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var users []*model.User
	for cursor.Next(context.TODO()) {
		var user model.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
