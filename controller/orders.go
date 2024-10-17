package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"plastiqu_co/model"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var orderCollection *mongo.Collection

// Initialize the order collection
func InitOrderController(client *mongo.Client) {
	orderCollection = client.Database("your_database").Collection("orders") // Ganti "your_database" dengan nama database Anda
}

// AddOrder untuk menambahkan pesanan baru
func AddOrder(w http.ResponseWriter, r *http.Request) {
	var order model.Orders

	// Decode JSON request body
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate a unique order code (for example: "ORD-USERID-TIMESTAMP")
	order.OrderCode = generateOrderCode(order.UserID)

	// Set default status based on payment method
	if order.PaymentMethod == "transfer" {
		order.Status = "Belum Bayar"
	} else if order.PaymentMethod == "COD" {
		order.Status = "Menunggu Konfirmasi"
	}

	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	// Insert order into the collection
	if _, err := orderCollection.InsertOne(context.TODO(), order); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

// GetAllOrders untuk mendapatkan semua pesanan
func GetAllOrders(w http.ResponseWriter, r *http.Request) {
	var orders []model.Orders

	cursor, err := orderCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var order model.Orders
		if err := cursor.Decode(&order); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		orders = append(orders, order)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

// GetOrderByID untuk mendapatkan pesanan berdasarkan ID
func GetOrderByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"]) // Mengubah ID ke ObjectID
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var order model.Orders
	err = orderCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&order)
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

// UpdateOrder untuk memperbarui pesanan berdasarkan ID
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"]) // Mengubah ID ke ObjectID
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var updatedOrder model.Orders

	// Decode JSON request body
	if err := json.NewDecoder(r.Body).Decode(&updatedOrder); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the order in the collection
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updatedOrder}

	result, err := orderCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil || result.MatchedCount == 0 {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedOrder)
}

// DeleteOrder untuk menghapus pesanan berdasarkan ID
func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"]) // Mengubah ID ke ObjectID
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Delete the order from the collection
	result, err := orderCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil || result.DeletedCount == 0 {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent) // Status 204 No Content
}

// generateOrderCode untuk menghasilkan kode pemesanan yang unik
func generateOrderCode(userID primitive.ObjectID) string {
	// Contoh pembuatan kode: "ORD-USERID-TIMESTAMP"
	return fmt.Sprintf("ORD-%s-%d", userID.Hex(), time.Now().UnixNano())
}

// AdvanceOrderStatus untuk mengubah status pesanan ke status berikutnya
func AdvanceOrderStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"]) // Mengubah ID ke ObjectID
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var order model.Orders
	if err := orderCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&order); err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	// Tentukan status berikutnya berdasarkan status saat ini
	var newStatus string
	switch order.Status {
	case "Belum Bayar":
		newStatus = "Sudah Bayar"
	case "Sudah Bayar":
		newStatus = "Diproses"
	case "Diproses":
		newStatus = "Dikemas"
	case "Dikemas":
		newStatus = "Dikirim"
	case "Dikirim":
		newStatus = "Selesai"
	default:
		http.Error(w, "No further status transition available", http.StatusBadRequest)
		return
	}

	// Update status
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": newStatus}}

	result, err := orderCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil || result.MatchedCount == 0 {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newStatus)
}
