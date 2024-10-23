package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Menu struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"` // Tambahkan ID unik
	Name          string             `json:"name" bson:"name"`
	Price         int                `json:"price" bson:"price"`
	OriginalPrice int                `json:"originalPrice" bson:"originalPrice"`
	Rating        float64            `json:"rating" bson:"rating"`
	Sold          int                `json:"sold" bson:"sold"`
	Image         string             `json:"image" bson:"image"`
	CreatedAt     time.Time          `json:"createdAt" bson:"createdAt"` // Tambahkan tanggal pembuatan
	UpdatedAt     time.Time          `json:"updatedAt" bson:"updatedAt"` // Tambahkan tanggal pembaruan
}
