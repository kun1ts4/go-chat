package middleware

import (
	"context"
	"go-chat/internal/services"
	"net/http"
)

func Auth(s *services.TokenService) func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header required", http.StatusUnauthorized)
				return
			}

			tokenString := authHeader[len("Bearer "):]
			username, err := s.ValidateToken(tokenString)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "user", username)
			handler.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
