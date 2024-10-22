package auth

import (
	"context"
	"encoding/json"
	"plastiqu_co/config"
	"plastiqu_co/model"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func LoginUsers(w http.ResponseWriter, r *http.Request) {
	var credentials model.Users
	var responseMessage string

	// Decode the request body into the credentials struct
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		responseMessage = "Invalid request payload"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   responseMessage,
			"message": "The JSON request body could not be decoded. Please check the structure of your request.",
		})
		return
	}

	// Check if email or password is empty
	if credentials.Email == "" || credentials.Password == "" {
		responseMessage = "Email and password are required"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   responseMessage,
			"message": "Please provide both email and password.",
		})
		return
	}

	// Setup a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get the MongoDB collection
	collection := config.Mongoconn.Collection("users")

	// Find the user by email
	var user model.Users
	err = collection.FindOne(ctx, bson.M{"email": credentials.Email}).Decode(&user)
	if err != nil {
		responseMessage = "User not found"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   responseMessage,
			"message": "The email is not registered.",
		})
		return
	}

	// Compare the provided password with the hashed password stored in the database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		responseMessage = "Invalid password"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   responseMessage,
			"message": "The password you entered is incorrect.",
		})
		return
	}

	// Send the response with user details if login is successful
	response := map[string]interface{}{
		"message":  "Login successful",
		"user_id":  user.ID.Hex(),
		"email":    user.Email,
		"username" : user.Username,
		"role":     user.Role,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
