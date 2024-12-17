package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Be2Bag/example/model"
	commonPorts "github.com/Be2Bag/example/module/common/ports"
	"github.com/Be2Bag/example/module/register/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RegisterRepository struct {
	collection *mongo.Collection
	sharedRepo commonPorts.SharedRepository
}

func NewRegisterRepository(db *mongo.Database, sharedRepo commonPorts.SharedRepository) ports.RegisterRepository {
	collection := db.Collection("users")
	return &RegisterRepository{
		collection: collection,
		sharedRepo: sharedRepo,
	}
}

func (repository *RegisterRepository) CreateUser(user *model.User) error {

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	_, err := repository.collection.InsertOne(context.TODO(), user)
	return err
}

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

func (repository *RegisterRepository) GetUserByID(user_id string) (*model.User, error) {

	var user model.User
	err := repository.collection.FindOne(context.TODO(), bson.M{"user_id": user_id}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (repository *RegisterRepository) UpdateUser(user *model.User) error {

	user.UpdatedAt = time.Now()
	_, err := repository.collection.UpdateOne(context.TODO(), bson.M{"user_id": user.UserID}, bson.M{"$set": user})
	if err != nil {
		return err
	}
	return nil
}

func (repository *RegisterRepository) DeleteUser(user_id string) error {
	_, err := repository.collection.DeleteOne(context.TODO(), bson.M{"user_id": user_id})
	if err != nil {
		return err
	}
	return nil
}
