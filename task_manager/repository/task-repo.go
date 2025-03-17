package repository

import (
	"task_manager/database"
	"task_manager/models"
)

// GetAllTasks retrieves tasks with optional filtering and pagination
func GetAllTasks(completed int, page int, pageSize int) ([]models.Task, error) {
	var tasks []models.Task
	offset := (page - 1) * pageSize
	query := database.DB.Model(&models.Task{})

	// Apply the filter for completed status
	if completed != -1 {
		query = query.Where("completed = ?", completed)
	}

	// Fetch tasks with pagination
	err := query.Offset(offset).Limit(pageSize).Find(&tasks).Error
	return tasks, err
}

// CreateTask adds a new task to the database
func CreateTask(task *models.Task) error {
	return database.DB.Create(task).Error
}

// UpdateTask updates an existing task's details
func UpdateTask(id uint, task *models.Task) (models.Task, error) {
	var updatedTask models.Task
	err := database.DB.Model(&models.Task{}).Where("id = ?", id).Updates(task).First(&updatedTask).Error
	if err != nil {
		return updatedTask, err
	}
	return updatedTask, nil
}

// MarkTaskAsCompleted updates the task's status to completed
func MarkTaskAsCompleted(id uint) error {
	var task models.Task
	if err := database.DB.First(&task, id).Error; err != nil {
		return err
	}
	task.Completed = true
	return database.DB.Save(&task).Error
}

// DeleteTask deletes a task from the database
func DeleteTask(id uint) error {
	return database.DB.Delete(&models.Task{}, id).Error
}

