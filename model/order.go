package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Order struct untuk menyimpan informasi pesanan
type Order struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	ProductID string             `json:"product_id" bson:"product_id"` // ID produk yang dipesan
	Quantity  int                `json:"quantity" bson:"quantity"`       // Jumlah produk yang dipesan
	Date      time.Time          `json:"date" bson:"date"`               // Tanggal pemesanan
}