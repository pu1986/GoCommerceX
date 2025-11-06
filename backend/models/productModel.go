package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Price       float64            `bson:"price" json:"price"`
	Category    string             `bson:"category" json:"category"`
	Image       string             `bson:"image" json:"image"`
	CreatedBy   primitive.ObjectID `bson:"createdBy,omitempty" json:"createdBy,omitempty"`
}
