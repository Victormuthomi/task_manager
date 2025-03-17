package repository

import (
	"task_manager/database"
	"task_manager/models"
)

// CreateTask inserts a new task into the database
func CreateTask(task *models.Task) error {
	return database.DB.Create(task).Error
}

// GetTasks retrieves all tasks from the database
func GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := database.DB.Find(&tasks).Error
	return tasks, err
}

// GetTaskByID retrieves a specific task by ID
func GetTaskByID(id uint) (models.Task, error) {
	var task models.Task
	err := database.DB.First(&task, id).Error
	return task, err
}

// UpdateTask modifies an existing task
func UpdateTask(task *models.Task) error {
	return database.DB.Save(task).Error
}

// DeleteTask removes a task from the database
func DeleteTask(id uint) error {
	return database.DB.Delete(&models.Task{}, id).Error
}

