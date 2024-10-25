package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Banner struct untuk menyimpan informasi tentang banner promosi
type Banner struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `bson:"title,omitempty" json:"title,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	ImageURL    string             `bson:"image_url,omitempty" json:"image_url,omitempty"`
	Active      bool               `bson:"active" json:"active"`
}
