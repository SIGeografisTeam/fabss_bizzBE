package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName       string             `bson:"first_name,omitempty" json:"first_name,omitempty"`
	LastName        string             `bson:"last_name,omitempty" json:"last_name,omitempty"`
	Email           string             `bson:"email,omitempty" json:"email,omitempty"`
	Role            string             `bson:"role,omitempty" json:"role,omitempty"` // Default "user"
	Phone           string             `bson:"phone,omitempty" json:"phone,omitempty"`
	Address         string             `bson:"address,omitempty" json:"address,omitempty"`
	Password        string             `bson:"password,omitempty" json:"password,omitempty"`
	ConfirmPassword string             `bson:"confirm_password,omitempty" json:"confirm_password,omitempty"` // Nullable
	Image           string             `bson:"image,omitempty" json:"image,omitempty"` // Nullable, URL or base64 for image
	CreatedAt       time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt       time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
