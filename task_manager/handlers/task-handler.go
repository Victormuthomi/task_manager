package handlers

import (
	"encoding/json"
	"task_manager/models"
	"task_manager/repository"
	"net/http"
	"strconv"
)

// GetTasks retrieves all tasks, optionally filtered by completion status
func GetTasks(w http.ResponseWriter, r *http.Request) {
	completedParam := r.URL.Query().Get("completed")
	pageParam := r.URL.Query().Get("page")
	pageSizeParam := r.URL.Query().Get("page_size")

	var completed int
	if completedParam == "true" {
		completed = 1
	} else if completedParam == "false" {
		completed = 0
	} else {
		completed = -1
	}

	page, err := strconv.Atoi(pageParam)
	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeParam)
	if err != nil {
		pageSize = 10
	}

	tasks, err := repository.GetAllTasks(completed, page, pageSize)
	if err != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// CreateTask adds a new task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := repository.CreateTask(&task); err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// UpdateTask updates an existing task
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	updatedTask, err := repository.UpdateTask(uint(taskID), &task) // Capture both returned values
	if err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTask) // Return the updated task
}

// MarkTaskAsCompleted marks a task as completed
func MarkTaskAsCompleted(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	err = repository.MarkTaskAsCompleted(uint(taskID))
	if err != nil {
		http.Error(w, "Failed to mark task as completed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteTask marks a task as deleted
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	err = repository.DeleteTask(uint(taskID))
	if err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

