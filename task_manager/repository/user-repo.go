package repository

import (
	"task_manager/models"
	"task_manager/database"
	"errors"
)

// CreateUser adds a new user to the database
func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

// GetUserByUsername retrieves a user by their username
func GetUserByUsername(username string) (models.User, error) {
	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return user, errors.New("user not found")
	}
	return user, nil
}

