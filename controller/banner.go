package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"plastiqu_co/config"
	"plastiqu_co/model"
)

// CreateBanner untuk menambahkan banner baru
func CreateBanner(w http.ResponseWriter, r *http.Request) {
	var banner model.Banner

	// Decode JSON request body ke struct Banner
	if err := json.NewDecoder(r.Body).Decode(&banner); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Invalid request payload",
			"message": "The JSON request body could not be decoded. Please check the structure of your request.",
		})
		return
	}

	banner.ID = primitive.NewObjectID()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := config.Mongoconn.Collection("banners")
	_, err := collection.InsertOne(ctx, banner)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Failed to create banner",
			"message": "An error occurred while creating the banner.",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Banner created successfully",
		"banner":  banner,
	})
}

// GetBanners untuk mendapatkan semua banner
func GetBanners(w http.ResponseWriter, r *http.Request) {
	var banners []model.Banner

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := config.Mongoconn.Collection("banners")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Failed to retrieve banners",
			"message": "An error occurred while retrieving banners.",
		})
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var banner model.Banner
		if err := cursor.Decode(&banner); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error":   "Failed to decode banner",
				"message": "An error occurred while decoding banner data.",
			})
			return
		}
		banners = append(banners, banner)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(banners)
}

// GetBannerByID untuk mengambil banner berdasarkan ID
func GetBannerByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Invalid banner ID",
			"message": "The provided banner ID is not valid.",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var banner model.Banner
	collection := config.Mongoconn.Collection("banners")
	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&banner)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Banner not found",
			"message": "The banner you are looking for does not exist.",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(banner)
}

// UpdateBanner untuk memperbarui banner berdasarkan ID
func UpdateBanner(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Invalid banner ID",
			"message": "The provided banner ID is not valid.",
		})
		return
	}

	var banner model.Banner
	if err := json.NewDecoder(r.Body).Decode(&banner); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Invalid request payload",
			"message": "The JSON request body could not be decoded. Please check the structure of your request.",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := config.Mongoconn.Collection("banners")
	filter := bson.M{"_id": id}
	update := bson.M{"$set": banner}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil || result.MatchedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Banner not found",
			"message": "The banner you are trying to update does not exist.",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Banner updated successfully",
	})
}

// DeleteBanner untuk menghapus banner berdasarkan ID
func DeleteBanner(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Invalid banner ID",
			"message": "The provided banner ID is not valid.",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := config.Mongoconn.Collection("banners")
	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil || result.DeletedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Banner not found",
			"message": "The banner you are trying to delete does not exist.",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Banner deleted successfully",
	})
}
