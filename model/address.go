package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Address represents an address model with the required fields
type Address struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FullName    string             `bson:"full_name" json:"full_name"`
	Phone       string             `bson:"phone" json:"phone"`
	Province    string             `bson:"province" json:"province"`
	City        string             `bson:"city" json:"city"`
	Subdistrict string             `bson:"subdistrict" json:"subdistrict"`
	PostalCode  string             `bson:"postal_code" json:"postal_code"`
	StreetName  string             `bson:"street_name" json:"street_name"`
	Building    string             `bson:"building,omitempty" json:"building,omitempty"`      // Optional field
	HouseNumber string             `bson:"house_number,omitempty" json:"house_number,omitempty"` // Optional field
	Label       string             `bson:"label" json:"label"` // e.g., "kantor" or "rumah"
}
