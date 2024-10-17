package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Orders struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID        primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`       // Foreign Key ke tabel Users
	ProductID     primitive.ObjectID `bson:"product_id,omitempty" json:"product_id,omitempty"` // Foreign Key ke tabel Product
	Fullname      string             `bson:"fullname,omitempty" json:"fullname,omitempty"`
	Phone         string             `bson:"phone,omitempty" json:"phone,omitempty"`
	Address       string             `bson:"address,omitempty" json:"address,omitempty"`
	ProductName   string             `bson:"product_name,omitempty" json:"product_name,omitempty"`
	Variant       string             `bson:"variant,omitempty" json:"variant,omitempty"` // Nullable
	Amount        int                `bson:"amount,omitempty" json:"amount,omitempty"`   // Jumlah produk
	Price         float64            `bson:"price,omitempty" json:"price,omitempty"`
	TotalPrice    float64            `bson:"total_price,omitempty" json:"total_price,omitempty"` // Total keseluruhan
	Status        string             `bson:"status,omitempty" json:"status,omitempty"`           // Menunggu Pembayaran, Diproses, Dikemas, Dikirim, Selesai, Dibatalkan
	PaymentMethod string             `bson:"payment_method,omitempty" json:"payment_method,omitempty"` // "transfer" atau "COD"
	PaymentProof  string             `bson:"payment_proof,omitempty" json:"payment_proof,omitempty"`   // URL Bukti Pembayaran (untuk transfer)
	CreatedAt     time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt     time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
