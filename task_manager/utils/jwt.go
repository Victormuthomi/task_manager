package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

var jwtSecret []byte

// Initialize the JWT secret from environment variables
func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load the JWT_SECRET from environment variables
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET environment variable is required")
	}
	jwtSecret = []byte(secret)
}

// GenerateJWT generates a JWT token for a given user ID
func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 72).Unix(), // Token expires in 72 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		log.Println("Failed to generate JWT:", err)
		return "", err
	}

	return signedToken, nil
}

// ValidateJWT validates a JWT token and returns the user ID if valid
func ValidateJWT(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		log.Println("Failed to parse JWT:", err)
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		log.Println("Invalid JWT claims")
		return 0, fmt.Errorf("invalid token claims")
	}

	userID, ok := claims["sub"].(float64)
	if !ok {
		log.Println("Failed to parse user ID from JWT")
		return 0, fmt.Errorf("failed to parse user ID")
	}

	return uint(userID), nil
}

