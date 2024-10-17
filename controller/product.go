package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"plastiqu_co/model" // Ganti dengan path ke model Anda
)

var productCollection *mongo.Collection

// Initialize the product collection
func InitProductController(client *mongo.Client) {
	productCollection = client.Database("your_database").Collection("products") // Ganti "your_database" dengan nama database Anda
}

// AddProduct untuk menambahkan produk baru
func AddProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product

	// Decode JSON request body
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert product into the collection
	if _, err := productCollection.InsertOne(context.TODO(), product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// GetAllProducts untuk mendapatkan semua produk
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	var products []model.Product

	cursor, err := productCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var product model.Product
		if err := cursor.Decode(&product); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// GetProductByID untuk mendapatkan produk berdasarkan ID
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var product model.Product
	err := productCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}
// UpdateProduct untuk memperbarui produk berdasarkan ID
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var updatedProduct model.Product

	// Decode JSON request body
	if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the product in the collection
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updatedProduct}

	result, err := productCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil || result.MatchedCount == 0 {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedProduct)
}

// DeleteProduct untuk menghapus produk berdasarkan ID
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	// Delete the product from the collection
	result, err := productCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil || result.DeletedCount == 0 {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent) // Status 204 No Content
}