package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CategoryName string             `bson:"name_category" json:"name_category"`
}
