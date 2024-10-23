package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Feedback struct to represent a product review
type Feedback struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	ProductID string             `json:"product_id" bson:"product_id"` // ID produk yang di-review
	UserID    string             `json:"user_id" bson:"user_id"`       // ID pengguna yang memberikan feedback
	Rating    int                `json:"rating" bson:"rating"`         // Rating dari 1-5
	Review    string             `json:"review" bson:"review"`         // Ulasan atau review
	Date      time.Time          `json:"date" bson:"date"`             // Tanggal dan waktu ulasan diberikan
	Response  *Response          `json:"response" bson:"response"`     // Balasan dari pemilik toko atau admin
}

type Response struct {
	UserID    string    `json:"user_id" bson:"user_id"` // ID pengguna yang memberikan balasan (admin/toko)
	Response  string    `json:"response" bson:"response"` // Isi balasan
	Date      time.Time `json:"date" bson:"date"`        // Tanggal dan waktu balasan diberikan
}

// WebsiteReview struct to represent a review for the website
type WebsiteReview struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    string    `json:"user_id" bson:"user_id"` // ID pengguna yang memberikan balasan (admin/toko)
	Rating    int                `bson:"rating" json:"rating"`
	Review    string             `bson:"review" json:"review"`
	Date      primitive.DateTime `bson:"date" json:"date"`
}

// MonthlySales struct untuk menyimpan data penjualan bulanan
type MonthlySales struct {
	ProductID  string `json:"product_id" bson:"product_id"`
	TotalSales int    `json:"total_sales" bson:"total_sales"`
}

// BestSeller struct untuk menyimpan data best seller
type BestSeller struct {
	ProductID string `json:"product_id" bson:"product_id"`
	Sales     int    `json:"sales" bson:"sales"`
	Clicks    int    `json:"clicks" bson:"clicks"` // Jumlah klik pada produk
}

// Session struct to represent a user session for tracking
type Session struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	IP       string             `bson:"ip" json:"ip"`
	Duration int                `bson:"duration" json:"duration"` // Duration in seconds
	Date     primitive.DateTime `bson:"date" json:"date"`
}

// AbandonedCart struct to represent abandoned cart tracking
type AbandonedCart struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    string             `bson:"user_id" json:"user_id"`
	ProductIDs []string           `bson:"product_ids" json:"product_ids"`
	Status    string             `bson:"status" json:"status"` // e.g., "abandoned"
	Date      primitive.DateTime `bson:"date" json:"date"`
}

// Cancellation struct to represent order cancellation
type Cancellation struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OrderID string             `bson:"order_id" json:"order_id"`
	Reason  string             `bson:"reason" json:"reason"`
	Date    primitive.DateTime `bson:"date" json:"date"`
}
