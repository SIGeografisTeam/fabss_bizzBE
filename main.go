package main

import (
	"fmt"
	"log"
	"net/http"

	"plastiqu_co/config"
	"plastiqu_co/routes"

	"github.com/rs/cors"
)

func main() {
	connectdb := config.Mongoconn

	if config.ErrorMongoconn != nil {
		fmt.Println("Failed to connect to MongoDB:", config.ErrorMongoconn)
		return
	}

	// Check if the connection is successful
	if connectdb != nil {
		fmt.Println("Successfully connected to MongoDB!")
	} else {
		fmt.Println("MongoDB connection is nil")
	}

	router := routes.InitializeRoutes()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            true,
	})

	handler := c.Handler(router)
	// Initialize the router from the routes package

	fmt.Println("Server is running on http://localhost:3600")
	log.Fatal(http.ListenAndServe(":3600", handler))
}
