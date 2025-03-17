package routes

import (
	"task_manager/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterRoutes sets up all API routes
func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()

	// Task Routes
	router.HandleFunc("/tasks", handlers.CreateTaskHandler).Methods("POST")
	router.HandleFunc("/tasks", handlers.GetTasksHandler).Methods("GET")
	router.HandleFunc("/tasks/{id}", handlers.GetTaskByIDHandler).Methods("GET")
	router.HandleFunc("/tasks/{id}", handlers.UpdateTaskHandler).Methods("PUT")
	router.HandleFunc("/tasks/{id}", handlers.DeleteTaskHandler).Methods("DELETE")

	// Health Check Route
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is running"))
	}).Methods("GET")

	return router
}

