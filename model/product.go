package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Variant struct untuk menyimpan informasi tentang varian produk
type Variant struct {
	Name  string  `bson:"name,omitempty" json:"name,omitempty"` // Nama varian
	Stock int     `bson:"stock,omitempty" json:"stock,omitempty"` // Stok untuk varian ini
	Price float64 `bson:"price,omitempty" json:"price,omitempty"` // Harga untuk varian ini
}

// Product struct untuk menyimpan informasi tentang produk
type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CategoryID  primitive.ObjectID `bson:"category_id,omitempty" json:"category_id,omitempty"` // Foreign Key ke tabel Category
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Variants    []Variant          `bson:"variants,omitempty" json:"variants,omitempty"` // Daftar varian produk (bisa kosong)
	TotalStock  int                `bson:"total_stock,omitempty" json:"total_stock,omitempty"` // Stok total (akumulasi dari semua varian atau stok produk itu sendiri)
	PriceRange  struct {
		Min float64 `bson:"min,omitempty" json:"min,omitempty"` // Harga terendah
		Max float64 `bson:"max,omitempty" json:"max,omitempty"` // Harga tertinggi
	} `bson:"price_range,omitempty" json:"price_range,omitempty"` // Rentang harga
	Stock       int                `bson:"stock,omitempty" json:"stock,omitempty"` // Stok produk jika tidak ada varian
	Price       float64            `bson:"price,omitempty" json:"price,omitempty"` // Harga produk jika tidak ada varian
	Image       string             `bson:"image,omitempty" json:"image,omitempty"`
}
