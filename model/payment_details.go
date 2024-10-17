package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentDetails struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Bank         string             `bson:"bank,omitempty" json:"bank,omitempty"`               // Nama Bank
	AccountName  string             `bson:"account_name,omitempty" json:"account_name,omitempty"` // Nama Pemilik Rekening
	AccountNumber string            `bson:"account_number,omitempty" json:"account_number,omitempty"` // Nomor Rekening
}
