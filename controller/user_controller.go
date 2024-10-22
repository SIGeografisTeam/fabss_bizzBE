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
	"golang.org/x/crypto/bcrypt"
)

// UpdateUserProfile allows a user to update their own profile
func UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(string) // assuming you have middleware to get user_id from the JWT
	var updateData model.Users

	// Decode the request body into updateData struct
	err := json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Invalid request payload",
			"message": "The JSON request body could not be decoded.",
		})
		return
	}

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

	// Prepare the update fields
	updateFields := bson.M{
		"username": updateData.Username,
		"phone":      updateData.Phone,
		"image":      updateData.Image, // Nullable
		"updated_at": time.Now(),
	}

	// Update the user's profile in the database
	collection := config.Mongoconn.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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
		"message": "Profile updated successfully",
	})
}

// ChangeUserPassword allows users to change their own password
func ChangeUserPassword(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(string)
	var requestData struct {
		CurrentPassword string `json:"current_password"`
		NewPassword     string `json:"new_password"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Invalid request payload",
			"message": "The JSON request body could not be decoded.",
		})
		return
	}

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

	collection := config.Mongoconn.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get the user from the database
	var user model.Users
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "User not found",
			"message": "The user with the specified ID does not exist.",
		})
		return
	}

	// Check if the current password is correct
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestData.CurrentPassword))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Incorrect current password",
			"message": "The current password you provided is incorrect.",
		})
		return
	}

	// Hash the new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestData.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Failed to hash password",
			"message": "An error occurred while hashing the password.",
		})
		return
	}

	// Update the user's password
	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": bson.M{"password": string(hashedPassword), "updated_at": time.Now()}})
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Failed to update password",
			"message": "An error occurred while updating the password.",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Password updated successfully",
	})
}
