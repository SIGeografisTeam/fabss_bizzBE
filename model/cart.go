package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Cart struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"` // Foreign Key ke tabel Users
	ProductID primitive.ObjectID `bson:"product_id,omitempty" json:"product_id,omitempty"` // Foreign Key ke tabel Product
	Variant   string             `bson:"variant,omitempty" json:"variant,omitempty"` // Nullable
	Amount    int                `bson:"amount,omitempty" json:"amount,omitempty"`   // Jumlah produk
	CreatedAt time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
