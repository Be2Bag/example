package repository

import (
	"github.com/Be2Bag/example/model"
	commonPorts "github.com/Be2Bag/example/module/common/ports"
	"github.com/Be2Bag/example/module/session/ports"
	"go.mongodb.org/mongo-driver/mongo"
)

type SessionRepository struct {
	collection *mongo.Collection
	sharedRepo commonPorts.SharedRepository
}

func NewSessionRepository(db *mongo.Database, sharedRepo commonPorts.SharedRepository) ports.SessionRepository {
	collection := db.Collection("users")
	return &SessionRepository{
		collection: collection,
		sharedRepo: sharedRepo,
	}
}

func (r *SessionRepository) GetUserByEmail(email string) (*model.Staff, error) {
	return r.sharedRepo.GetStaffByEmail(email)
}
