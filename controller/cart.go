package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"plastiqu_co/model"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var cartCollection *mongo.Collection

// Initialize the cart collection
func InitCartController(client *mongo.Client) {
	cartCollection = client.Database("your_database").Collection("carts") // Ganti "your_database" dengan nama database Anda
}

// AddToCart untuk menambahkan produk ke keranjang
func AddToCart(w http.ResponseWriter, r *http.Request) {
	var cart model.Cart

	// Decode JSON request body
	if err := json.NewDecoder(r.Body).Decode(&cart); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cart.CreatedAt = time.Now()
	cart.UpdatedAt = time.Now()

	// Insert cart item into the collection
	if _, err := cartCollection.InsertOne(context.TODO(), cart); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cart)
}

// GetCartItems untuk mendapatkan semua item di keranjang untuk pengguna tertentu
func GetCartItems(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := primitive.ObjectIDFromHex(params["user_id"]) // Mengubah user_id ke ObjectID
	if err != nil {
		http.Error(w, "Invalid user ID format", http.StatusBadRequest)
		return
	}

	var carts []model.Cart

	cursor, err := cartCollection.Find(context.TODO(), bson.M{"user_id": userID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var cart model.Cart
		if err := cursor.Decode(&cart); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		carts = append(carts, cart)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(carts)
}

// UpdateCartItem untuk memperbarui item di keranjang berdasarkan ID
func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"]) // Mengubah ID ke ObjectID
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var updatedCart model.Cart

	// Decode JSON request body
	if err := json.NewDecoder(r.Body).Decode(&updatedCart); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the cart item in the collection
	filter := bson.M{"_id": id}
	updatedCart.UpdatedAt = time.Now() // Update the timestamp

	update := bson.M{"$set": updatedCart}

	result, err := cartCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil || result.MatchedCount == 0 {
		http.Error(w, "Cart item not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedCart)
}

// RemoveFromCart untuk menghapus item dari keranjang berdasarkan ID
func RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"]) // Mengubah ID ke ObjectID
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Delete the cart item from the collection
	result, err := cartCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil || result.DeletedCount == 0 {
		http.Error(w, "Cart item not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent) // Status 204 No Content
}
