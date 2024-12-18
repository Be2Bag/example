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
	collection := db.Collection("staff")
	return &RegisterRepository{
		collection: collection,
		sharedRepo: sharedRepo,
	}
}

func (repository *RegisterRepository) CreateStaff(staff *model.Staff) error {

	staff.CreatedAt = time.Now()
	staff.UpdatedAt = time.Now()
	_, err := repository.collection.InsertOne(context.TODO(), staff)
	return err
}

func (repository *RegisterRepository) GetStaffByEmail(email string) (*model.Staff, error) {
	var staff model.Staff
	err := repository.collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&staff)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &staff, nil
}

func (repository *RegisterRepository) GetStaffs() ([]*model.Staff, error) {
	cursor, err := repository.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var staffs []*model.Staff
	for cursor.Next(context.TODO()) {
		var staff model.Staff
		if err := cursor.Decode(&staff); err != nil {
			return nil, err
		}
		staffs = append(staffs, &staff)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return staffs, nil
}

func (repository *RegisterRepository) GetStaffByID(staff_id string) (*model.Staff, error) {

	var staff model.Staff
	err := repository.collection.FindOne(context.TODO(), bson.M{"user_id": staff_id}).Decode(&staff)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &staff, nil
}

func (repository *RegisterRepository) UpdateStaff(staff *model.Staff) error {

	staff.UpdatedAt = time.Now()
	_, err := repository.collection.UpdateOne(context.TODO(), bson.M{"user_id": staff.UserID}, bson.M{"$set": staff})
	if err != nil {
		return err
	}
	return nil
}

func (repository *RegisterRepository) DeleteStaff(staff_id string) error {
	_, err := repository.collection.DeleteOne(context.TODO(), bson.M{"user_id": staff_id})
	if err != nil {
		return err
	}
	return nil
}
