package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Review struct untuk menyimpan ulasan produk
type Review struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // ID ulasan
	ProductID      primitive.ObjectID `bson:"product_id,omitempty" json:"product_id,omitempty"` // Foreign Key ke tabel Product
	UserID         primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"` // Foreign Key ke user yang memberikan ulasan
	Username       string             `bson:"username,omitempty" json:"username,omitempty"` // Nama pengguna yang memberikan ulasan
	Rating         int                `bson:"rating,omitempty" json:"rating,omitempty"` // Rating (nilai ulasan)
	ReviewText     string             `bson:"review_text,omitempty" json:"review_text,omitempty"` // Teks ulasan
	ReviewImage    string             `bson:"review_image,omitempty" json:"review_image,omitempty"` // Foto yang disertakan dalam ulasan (opsional)
	AdminResponse  string             `bson:"admin_response,omitempty" json:"admin_response,omitempty"` // Tanggapan admin (opsional)
	ResponseDate   time.Time          `bson:"response_date,omitempty" json:"response_date,omitempty"` // Waktu tanggapan admin (opsional)
	CreatedAt      time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"` // Tanggal ulasan dibuat
	UpdatedAt      time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"` // Tanggal ulasan terakhir diperbarui
}


// WebsiteReview struct untuk merepresentasikan ulasan tentang website
type WebsiteReview struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"` // ID ulasan
	UserID    string             `json:"user_id" bson:"user_id"`   // ID pengguna yang memberikan ulasan
	Rating    int                `bson:"rating" json:"rating"`     // Rating dari 1-5
	Review    string             `bson:"review" json:"review"`     // Ulasan
	Date      time.Time          `bson:"date" json:"date"`         // Tanggal dan waktu ulasan diberikan
}