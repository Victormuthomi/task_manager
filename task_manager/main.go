package main

import (
	"fmt"
	"task_manager/database"
	"task_manager/routes"
	"log"
	"net/http"
)

func main() {
	// Connect to the database
	database.ConnectDB()

	// Register API routes
	router := routes.RegisterRoutes()

	// Start the server
	port := ":8080"
	fmt.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}

