package models

import "gopkg.in/mgo.v2/bson"

const (
	// CollectionUser holds the name of the users collection
	CollectionUser = "users"
)

// User model
type User struct {
	Id       bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string        `json:"username" binding:"required" bson:"username"`
	Email    string        `json:"email" binding:"required" bson:"email"`
	Password string        `json:"password" binding:"required" bson:"password"`
}
