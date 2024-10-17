package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux" // Import gorilla/mux untuk menangani path parameters
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"plastiqu_co/config" // Ganti dengan path package config Anda
	"plastiqu_co/model"  // Ganti dengan path package model Anda
)

// CreateCategory handles the creation of a new category
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category model.Category

	// Decode the JSON request body into the category struct
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Invalid request payload",
			"message": "The JSON request body could not be decoded. Please check the structure of your request.",
		})
		return
	}

	// Create the category with current time
	category.ID = primitive.NewObjectID() // Generate a new ID
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := config.Mongoconn.Collection("categories")
	_, err = collection.InsertOne(ctx, category)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Failed to create category",
			"message": "An error occurred while creating the category.",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Category created successfully",
		"category": category,
	})
}

// GetCategories retrieves all categories
func GetCategories(w http.ResponseWriter, r *http.Request) {
	var categories []model.Category

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := config.Mongoconn.Collection("categories")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Failed to retrieve categories",
			"message": "An error occurred while retrieving categories.",
		})
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var category model.Category
		if err := cursor.Decode(&category); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error":   "Failed to decode category",
				"message": "An error occurred while decoding category data.",
			})
			return
		}
		categories = append(categories, category)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
}

// GetCategoryByID retrieves a category by its ID
func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"]) // Mengambil ID dari URL path

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Invalid category ID",
			"message": "The provided category ID is not valid.",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var category model.Category
	collection := config.Mongoconn.Collection("categories")
	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&category)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Category not found",
			"message": "The category you are looking for does not exist.",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(category)
}

// UpdateCategory updates an existing category
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"]) // Mengambil ID dari URL path

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Invalid category ID",
			"message": "The provided category ID is not valid.",
		})
		return
	}

	var category model.Category
	// Decode the JSON request body into the category struct
	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Invalid request payload",
			"message": "The JSON request body could not be decoded. Please check the structure of your request.",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := config.Mongoconn.Collection("categories")
	filter := bson.M{"_id": id} // Filter berdasarkan ID
	update := bson.M{"$set": category}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil || result.MatchedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Category not found",
			"message": "The category you are trying to update does not exist.",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Category updated successfully",
	})
}

// DeleteCategory deletes an existing category
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"]) // Mengambil ID dari URL path

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Invalid category ID",
			"message": "The provided category ID is not valid.",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := config.Mongoconn.Collection("categories")
	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil || result.DeletedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Category not found",
			"message": "The category you are trying to delete does not exist.",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Category deleted successfully",
	})
}
