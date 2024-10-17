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

// CreatePaymentDetails creates a new payment detail
func CreatePaymentDetails(w http.ResponseWriter, r *http.Request) {
	var payment model.PaymentDetails

	err := json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Invalid request payload",
			"message": "The JSON request body could not be decoded.",
		})
		return
	}

	payment.ID = primitive.NewObjectID() // Generate a new ID
	collection := config.Mongoconn.Collection("payment_details")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, payment)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Failed to create payment details",
			"message": "An error occurred while creating payment details.",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(payment)
}

// GetPaymentDetails retrieves all payment details
func GetPaymentDetails(w http.ResponseWriter, r *http.Request) {
	collection := config.Mongoconn.Collection("payment_details")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Failed to retrieve payment details",
			"message": "An error occurred while retrieving payment details.",
		})
		return
	}
	defer cursor.Close(ctx)

	var payments []model.PaymentDetails
	for cursor.Next(ctx) {
		var payment model.PaymentDetails
		if err := cursor.Decode(&payment); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error":   "Failed to decode payment detail",
				"message": "An error occurred while decoding payment details.",
			})
			return
		}
		payments = append(payments, payment)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payments)
}

// GetPaymentDetailByID retrieves a payment detail by ID
func GetPaymentDetailByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id") // Assume the ID is passed as a query parameter

	if id == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Missing payment detail ID",
			"message": "You must provide an ID to retrieve payment details.",
		})
		return
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Invalid payment detail ID",
			"message": "Payment detail ID format is incorrect.",
		})
		return
	}

	collection := config.Mongoconn.Collection("payment_details")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var payment model.PaymentDetails
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&payment)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Payment detail not found",
			"message": "No payment detail found with the specified ID.",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payment)
}

// UpdatePaymentDetails updates an existing payment detail
func UpdatePaymentDetails(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id") // Assume the ID is passed as a query parameter

	if id == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Missing payment detail ID",
			"message": "You must provide an ID to update payment details.",
		})
		return
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Invalid payment detail ID",
			"message": "Payment detail ID format is incorrect.",
		})
		return
	}

	var payment model.PaymentDetails
	err = json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Invalid request payload",
			"message": "The JSON request body could not be decoded.",
		})
		return
	}

	collection := config.Mongoconn.Collection("payment_details")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": payment})
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Failed to update payment details",
			"message": "An error occurred while updating payment details.",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Payment details updated successfully",
	})
}

// DeletePaymentDetails deletes a payment detail by ID
func DeletePaymentDetails(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id") // Assume the ID is passed as a query parameter

	if id == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Missing payment detail ID",
			"message": "You must provide an ID to delete payment details.",
		})
		return
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Invalid payment detail ID",
			"message": "Payment detail ID format is incorrect.",
		})
		return
	}

	collection := config.Mongoconn.Collection("payment_details")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Failed to delete payment details",
			"message": "An error occurred while deleting payment details.",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Payment details deleted successfully",
	})
}
