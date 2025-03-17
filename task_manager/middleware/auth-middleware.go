package middleware

import (
    "net/http"
    "strings"
    "task_manager/utils"
    "context" // Import context package
    "log"
)

// AuthMiddleware is a middleware that checks the JWT token in the request header
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Retrieve the token from the Authorization header
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
            return
        }

        // Remove "Bearer " from the start of the header value
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")

        // Validate the JWT token and get the user ID
        userID, err := utils.ValidateJWT(tokenString)
        if err != nil {
            http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
            return
        }

        // Store the user ID in the request context
        ctx := context.WithValue(r.Context(), "userID", userID)

        // Pass the request to the next handler with the updated context
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// GetUserIDFromContext retrieves the user ID from the request context
func GetUserIDFromContext(r *http.Request) uint {
    userID, ok := r.Context().Value("userID").(uint)
    if !ok {
        log.Println("Error: User ID not found in context")
        return 0
    }
    return userID
}

