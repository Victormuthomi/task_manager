package models

import "gorm.io/gorm"

// Task represents a task in the system
type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

