package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"plastiqu_co/model"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

	// Set the ID to a new ObjectID if it is not provided
	if product.ID.IsZero() {
		product.ID = primitive.NewObjectID()
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
	// Convert id string to ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	err = productCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&product)
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

	// Convert id string to ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Update the product in the collection
	filter := bson.M{"_id": objID}
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

	// Convert id string to ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Delete the product from the collection
	result, err := productCollection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil || result.DeletedCount == 0 {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent) // Status 204 No Content
}
