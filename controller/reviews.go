package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux" // Import gorilla/mux untuk menangani path parameters
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"plastiqu_co/config"
	"plastiqu_co/model"
)

// CreateReviewHandler untuk membuat ulasan baru
func CreateReviewHandler(w http.ResponseWriter, r *http.Request) {
	var review model.Review

	// Decode request body ke struct Review
	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Invalid request payload",
			"message": "The JSON request body could not be decoded. Please check the structure of your request.",
		})
		return
	}

	// Tentukan waktu saat ulasan dibuat
	review.CreatedAt = time.Now()
	review.UpdatedAt = time.Now()

	// Generate ObjectID baru untuk ulasan
	review.ID = primitive.NewObjectID()

	// Simpan ulasan di database
	collection := config.Mongoconn.Collection("reviews")
	_, err := collection.InsertOne(context.TODO(), review)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Failed to create review",
			"message": "An error occurred while creating the review.",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(review)
}

// GetReviewsHandler untuk mengambil semua ulasan produk
func GetReviewsHandler(w http.ResponseWriter, r *http.Request) {
	productID := mux.Vars(r)["product_id"]
	objID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Invalid product ID",
			"message": "The provided product ID is not valid.",
		})
		return
	}

	// Query untuk mengambil ulasan berdasarkan productID
	collection := config.Mongoconn.Collection("reviews")
	cursor, err := collection.Find(context.TODO(), bson.M{"product_id": objID})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Failed to fetch reviews",
			"message": "An error occurred while fetching the reviews.",
		})
		return
	}
	defer cursor.Close(context.TODO())

	var reviews []model.Review
	if err := cursor.All(context.TODO(), &reviews); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Failed to decode reviews",
			"message": "An error occurred while decoding review data.",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(reviews)
}

// UpdateReviewHandler untuk memperbarui ulasan
func UpdateReviewHandler(w http.ResponseWriter, r *http.Request) {
	reviewID := mux.Vars(r)["review_id"]
	objID, err := primitive.ObjectIDFromHex(reviewID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Invalid review ID",
			"message": "The provided review ID is not valid.",
		})
		return
	}

	var review model.Review
	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Invalid request payload",
			"message": "The JSON request body could not be decoded. Please check the structure of your request.",
		})
		return
	}

	// Tentukan waktu saat ulasan diperbarui
	review.UpdatedAt = time.Now()

	// Update ulasan di database
	collection := config.Mongoconn.Collection("reviews")
	_, err = collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		bson.M{
			"$set": bson.M{
				"review_text":  review.ReviewText,
				"rating":       review.Rating,
				"review_image": review.ReviewImage,
				"updated_at":   review.UpdatedAt,
			},
		},
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Failed to update review",
			"message": "An error occurred while updating the review.",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Review updated successfully",
	})
}

// DeleteReviewHandler untuk menghapus ulasan
func DeleteReviewHandler(w http.ResponseWriter, r *http.Request) {
	reviewID := mux.Vars(r)["review_id"]
	objID, err := primitive.ObjectIDFromHex(reviewID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Invalid review ID",
			"message": "The provided review ID is not valid.",
		})
		return
	}

	// Hapus ulasan dari database
	collection := config.Mongoconn.Collection("reviews")
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Failed to delete review",
			"message": "An error occurred while deleting the review.",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Review deleted successfully",
	})
}

// AdminRespondReviewHandler untuk menanggapi ulasan
func AdminRespondReviewHandler(w http.ResponseWriter, r *http.Request) {
	reviewID := mux.Vars(r)["review_id"]
	objID, err := primitive.ObjectIDFromHex(reviewID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Invalid review ID",
			"message": "The provided review ID is not valid.",
		})
		return
	}

	var response struct {
		AdminResponse string `json:"admin_response"`
	}

	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Invalid request payload",
			"message": "The JSON request body could not be decoded. Please check the structure of your request.",
		})
		return
	}

	// Tentukan waktu saat tanggapan diberikan
	responseDate := time.Now()

	// Update tanggapan admin di ulasan
	collection := config.Mongoconn.Collection("reviews")
	_, err = collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		bson.M{
			"$set": bson.M{
				"admin_response": response.AdminResponse,
				"response_date":  responseDate,
			},
		},
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Failed to respond to review",
			"message": "An error occurred while responding to the review.",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Admin responded successfully",
	})
}
