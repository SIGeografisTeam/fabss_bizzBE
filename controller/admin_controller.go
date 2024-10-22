package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"plastiqu_co/config"
	"plastiqu_co/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AdminUpdateUserProfile allows an admin to update any user's profile
func AdminUpdateUserProfile(w http.ResponseWriter, r *http.Request) {
	adminID := r.URL.Query().Get("admin_id") // ambil admin_id dari query parameter
	var updateData model.Users

	// Get the user ID from the URL
	userID := r.URL.Query().Get("user_id")

	// Convert userID string to ObjectID
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Invalid user ID",
			"message": "User ID format is incorrect.",
		})
		return
	}

	// Decode the request body into updateData struct
	err = json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Invalid request payload",
			"message": "The JSON request body could not be decoded.",
		})
		return
	}

	// Get admin user data to check role
	collection := config.Mongoconn.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var adminUser model.Users
	err = collection.FindOne(ctx, bson.M{"_id": adminID}).Decode(&adminUser)
	if err != nil || adminUser.Role != "admin" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Access denied",
			"message": "You do not have permission to perform this action.",
		})
		return
	}

	// Prepare the update fields
	updateFields := bson.M{
		"username": updateData.Username,
		"phone":      updateData.Phone,
		"image":      updateData.Image, // Nullable
		"role":       updateData.Role,  // Admin can update role
		"updated_at": time.Now(),
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": updateFields})
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Failed to update profile",
			"message": "An error occurred while updating the profile.",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User profile updated successfully",
	})
}

// AdminUpdateUserRole allows an admin to upgrade a user's role to admin
func AdminUpdateUserRole(w http.ResponseWriter, r *http.Request) {
	adminID := r.URL.Query().Get("admin_id") // ambil admin_id dari query parameter

	// Get the user ID from the URL
	userID := r.URL.Query().Get("user_id")
	
	// Convert userID string to ObjectID
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Invalid user ID",
			"message": "User ID format is incorrect.",
		})
		return
	}

	// Get admin user data to check role
	collection := config.Mongoconn.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var adminUser model.Users
	err = collection.FindOne(ctx, bson.M{"_id": adminID}).Decode(&adminUser)
	if err != nil || adminUser.Role != "admin" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Access denied",
			"message": "You do not have permission to perform this action.",
		})
		return
	}

	// Update the user's role to admin
	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": bson.M{"role": "admin", "updated_at": time.Now()}})
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Failed to update role",
			"message": "An error occurred while updating the user's role.",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User role updated to admin successfully",
	})
}
