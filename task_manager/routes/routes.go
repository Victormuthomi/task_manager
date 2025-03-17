package routes

import (
	"task_manager/handlers"
	"task_manager/middleware"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
	r.HandleFunc("/login", handlers.LoginUser).Methods("POST")

	// Task routes
	taskRoutes := r.PathPrefix("/tasks").Subrouter()
	taskRoutes.Use(middleware.AuthMiddleware)
	taskRoutes.HandleFunc("", handlers.GetTasks).Methods("GET")
	taskRoutes.HandleFunc("", handlers.CreateTask).Methods("POST")
	taskRoutes.HandleFunc("/{id:[0-9]+}", handlers.UpdateTask).Methods("PUT")
	taskRoutes.HandleFunc("/{id:[0-9]+}/complete", handlers.MarkTaskAsCompleted).Methods("PATCH")
	taskRoutes.HandleFunc("/{id:[0-9]+}/delete", handlers.DeleteTask).Methods("DELETE")
}

