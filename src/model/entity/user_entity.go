package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEntity struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstName,omitempty"`
	LastName  string             `json:"lastName,omitempty"`
	Email     string             `json:"email,omitempty"`
	Age       int8               `json:"age,omitempty"`
}
