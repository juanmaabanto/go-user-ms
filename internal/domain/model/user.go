package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserName    string             `bson:"userName"`
	FirstName   string             `bson:"firstName"`
	LastName    string             `bson:"lastName"`
	Email       string             `bson:"email"`
	PhoneNumber string             `bson:"phoneNumber"`
	BirthDate   time.Time          `json:"birthDate" bson:"birthDate"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	CreatedBy   string             `json:"createdBy" bson:"createdBy"`
	ModifiedAt  *time.Time         `json:"modifiedAt,omitempty" bson:"modifiedAt,omitempty"`
	ModifiedBy  *string            `json:"modifiedBy,omitempty" bson:"modifiedBy,omitempty"`
}

func (User) GetCollectionName() string {
	return "users"
}
