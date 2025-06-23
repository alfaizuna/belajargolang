package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"library-api/internal/models"
	"library-api/internal/repository"

	"github.com/golang-jwt/jwt/v5"
)

type key int

const userKey key = 0

func AuthMiddleware(userRepo *repository.UserRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			if header == "" || !strings.HasPrefix(header, "Bearer ") {
				http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
				return
			}
			tokenStr := strings.TrimPrefix(header, "Bearer ")
			secret := os.Getenv("JWT_SECRET")
			if secret == "" {
				secret = "secret"
			}
			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})
			if err != nil || !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}
			userIDFloat, ok := claims["user_id"].(float64)
			if !ok {
				http.Error(w, "Invalid user_id in token", http.StatusUnauthorized)
				return
			}
			userID := uint(userIDFloat)
			var user models.User
			err = userRepo.DB.First(&user, userID).Error
			if err != nil {
				http.Error(w, "User not found", http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), "user", &user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// RBAC middleware: hanya izinkan role tertentu
func RequireRole(role models.Role) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, ok := r.Context().Value("user").(*models.User)
			if !ok || user.Role != role {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
