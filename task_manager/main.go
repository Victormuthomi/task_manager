package main

import (
	"fmt"
	"log"
	"net/http"
	"task_manager/database"
	"task_manager/routes"
	"github.com/gorilla/mux"
)

func main() {
	// Connect to the database
	database.ConnectDB()

	// Create a new router
	router := mux.NewRouter()

	// Register API routes by passing the router to RegisterRoutes
	routes.RegisterRoutes(router)

	// Start the server
	port := ":8080"
	fmt.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}

