package auth

import (
	"encoding/json"
	"plastiqu_co/config"
	"plastiqu_co/model"
	"net/http"
	"time"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUsers(w http.ResponseWriter, r *http.Request) {
	var user model.Users
	var responseMessage string

	// Decode the JSON request body into the user struct
err := json.NewDecoder(r.Body).Decode(&user)
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

	// Validasi masing-masing field
	if user.FirstName == "" {
		responseMessage = "First name is required"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   responseMessage,
			"message": "Please provide a valid first name.",
		})
		return
	}

	if user.LastName == "" {
		responseMessage = "Last name is required"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   responseMessage,
			"message": "Please provide a valid last name.",
		})
		return
	}

	if user.Phone == "" {
		responseMessage = "Phone number is required"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   responseMessage,
			"message": "Please provide a valid phone number.",
		})
		return
	}

	if user.Email == "" {
		responseMessage = "Email is required"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   responseMessage,
			"message": "Please provide a valid email address.",
		})
		return
	}

	if user.Password == "" {
		responseMessage = "Password is required"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   responseMessage,
			"message": "Please provide a password.",
		})
		return
	}
	
	// Check if confirm_password exists in request body
	if user.ConfirmPassword == "" {
		responseMessage = "Confirm password is required"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   responseMessage,
			"message": "Please provide a confirm password.",
		})
		return
	}
	
	// Check if password and confirm password match
	if user.Password != user.ConfirmPassword {
		responseMessage = "Passwords do not match"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   responseMessage,
			"message": "The passwords you entered do not match.",
		})
		return
	}

	// Hash the user's password before saving it to the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		responseMessage = "Failed to hash password"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   responseMessage,
			"message": "An error occurred while hashing the password.",
		})
		return
	}
	user.Password = string(hashedPassword)

	user.Role = "user" // Set default role to "user"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	collection := config.Mongoconn.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, bson.M{
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"phone":      user.Phone,
		"email":      user.Email,
		"address":    user.Address,
		"role":       user.Role,
		"password":   user.Password,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	})

	if err != nil {
		responseMessage = "Failed to insert user"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   responseMessage,
			"message": "An error occurred while inserting the user into the database.",
		})
		return
	}

	response := map[string]interface{}{
		"message": "User registered successfully",
		"user_id": result.InsertedID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
