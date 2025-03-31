package models

import "go.mongodb.org/mongo-driver/v2/bson"

type User struct {
	Id       bson.ObjectID `json:"_id"`
	Username string        `json:"username"`
	Password string        `json:"password"`
	Email    string        `json:"email"`
}
