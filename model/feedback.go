package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Feedback struct untuk merepresentasikan ulasan produk
type Feedback struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`          // ID ulasan
	ProductID string             `json:"product_id" bson:"product_id"` // ID produk yang di-review
	UserID    string             `json:"user_id" bson:"user_id"`       // ID pengguna yang memberikan feedback
	Rating    int                `json:"rating" bson:"rating"`         // Rating dari 1-5
	Review    string             `json:"review" bson:"review"`         // Ulasan atau feedback
	Date      time.Time          `json:"date" bson:"date"`             // Tanggal dan waktu ulasan diberikan
	Response  *Response          `json:"response,omitempty" bson:"response,omitempty"` // Balasan dari pemilik toko atau admin
}

// Response struct untuk merepresentasikan balasan dari pemilik toko atau admin
type Response struct {
	UserID    string    `json:"user_id" bson:"user_id"` // ID pengguna yang memberikan balasan (admin/toko)
	Response  string    `json:"response" bson:"response"` // Isi balasan
	Date      time.Time `json:"date" bson:"date"`        // Tanggal dan waktu balasan diberikan
}

// MonthlySales struct untuk menyimpan data penjualan bulanan
type MonthlySales struct {
	ProductID  string `json:"product_id" bson:"product_id"` // ID produk
	TotalSales int    `json:"total_sales" bson:"total_sales"` // Total penjualan produk
}

// BestSeller struct untuk menyimpan data best seller
type BestSeller struct {
	ProductID string `json:"product_id" bson:"product_id"` // ID produk
	Sales     int    `json:"sales" bson:"sales"`             // Jumlah penjualan produk
	Clicks    int    `json:"clicks" bson:"clicks"`           // Jumlah klik pada produk
}

type FavoriteProduct struct {
    ProductID  primitive.ObjectID `bson:"_id" json:"product_id"`
    TotalClicks int               `bson:"totalClicks" json:"total_clicks"`
}

type TopSellingProduct struct {
    ProductID   primitive.ObjectID `bson:"_id" json:"product_id"`
    TotalSales  int                `bson:"totalSales" json:"total_sales"`
}


// Session struct untuk merepresentasikan sesi pengguna untuk pelacakan
type Session struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"` // ID sesi
	IP       string             `bson:"ip" json:"ip"`            // IP pengguna
	Duration int                `bson:"duration" json:"duration"` // Durasi sesi dalam detik
	Date     time.Time          `bson:"date" json:"date"`        // Tanggal dan waktu sesi
}

// AbandonedCart struct untuk merepresentasikan pelacakan keranjang yang ditinggalkan
type AbandonedCart struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"` // ID keranjang
	UserID    string             `bson:"user_id" json:"user_id"`   // ID pengguna
	ProductIDs []string           `bson:"product_ids" json:"product_ids"` // ID produk dalam keranjang
	Status    string             `bson:"status" json:"status"`     // Status keranjang, e.g., "abandoned"
	Date      time.Time          `bson:"date" json:"date"`        // Tanggal dan waktu keranjang ditinggalkan
}

// Cancellation struct untuk merepresentasikan pembatalan pesanan
type Cancellation struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"` // ID pembatalan
	OrderID string             `bson:"order_id" json:"order_id"` // ID pesanan yang dibatalkan
	Reason  string             `bson:"reason" json:"reason"`     // Alasan pembatalan
	Date    time.Time          `bson:"date" json:"date"`        // Tanggal dan waktu pembatalan
}
