package models

import "go.mongodb.org/mongo-driver/v2/bson"

type User struct {
	Id       bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	Username string        `bson:"username" json:"username" validate:"min=3,max=20"`
	Password string        `bson:"password" json:"password" validate:"min=6,max=20"`
	Email    string        `bson:"email" json:"email" validate:"email"`
}
