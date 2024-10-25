package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"plastiqu_co/config"
	"plastiqu_co/model"
)

// CreateAddress untuk menambahkan address baru
func CreateAddress(w http.ResponseWriter, r *http.Request) {
	var address model.Address
	if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	address.ID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := config.Mongoconn.Collection("addresses")
	_, err := collection.InsertOne(ctx, address)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create address"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(address)
}

// GetAddresses untuk mendapatkan semua address
func GetAddresses(w http.ResponseWriter, r *http.Request) {
	var addresses []model.Address

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := config.Mongoconn.Collection("addresses")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to retrieve addresses"})
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var address model.Address
		if err := cursor.Decode(&address); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to decode address data"})
			return
		}
		addresses = append(addresses, address)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(addresses)
}

// GetAddressByID untuk mendapatkan address berdasarkan ID
func GetAddressByID(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid address ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var address model.Address
	collection := config.Mongoconn.Collection("addresses")
	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&address)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Address not found"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(address)
}

// UpdateAddress untuk memperbarui address berdasarkan ID
func UpdateAddress(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid address ID"})
		return
	}

	var address model.Address
	if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := config.Mongoconn.Collection("addresses")
	update := bson.M{"$set": address}
	result, err := collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil || result.MatchedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Address not found"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Address updated successfully"})
}

// DeleteAddress untuk menghapus address berdasarkan ID
func DeleteAddress(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid address ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := config.Mongoconn.Collection("addresses")
	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil || result.DeletedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Address not found"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Address deleted successfully"})
}
