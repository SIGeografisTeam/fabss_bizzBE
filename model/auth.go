package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username       string             `bson:"username,omitempty" json:"username,omitempty"`
	Email           string             `bson:"email,omitempty" json:"email,omitempty"`
	Role            string             `bson:"role,omitempty" json:"role,omitempty"` // Default "user"
	Phone           string             `bson:"phone,omitempty" json:"phone,omitempty"`
	Password        string             `bson:"password,omitempty" json:"password,omitempty"`
	Image           string             `bson:"image,omitempty" json:"image,omitempty"` // Nullable, URL or base64 for image
	CreatedAt       time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt       time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
