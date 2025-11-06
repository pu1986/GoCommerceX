package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `bson:"name" json:"name" binding:"required"`
	Email    string             `bson:"email" json:"email" binding:"required"`
	Password string             `bson:"password" json:"password" binding:"required"`
	IsAdmin  bool               `bson:"isAdmin" json:"isAdmin"`
}
