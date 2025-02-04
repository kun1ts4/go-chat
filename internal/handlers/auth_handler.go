package handlers

import (
	"context"
	"encoding/json"
	"go-chat/internal/domain"
	"go-chat/internal/services"
	"io"
	"net/http"
)

type AuthHandler struct {
	userService  *services.UserService
	tokenService *services.TokenService
}

func NewAuthHandler(userService *services.UserService, service *services.TokenService) *AuthHandler {
	return &AuthHandler{
		userService:  userService,
		tokenService: service,
	}
}
func (a *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		user := domain.User{}
		err = json.Unmarshal(body, &user)
		if err != nil {
			http.Error(w, "Error parsing request body", http.StatusBadRequest)
			return
		}

		err = a.userService.Register(context.Background(), user)
		if err != nil {
			http.Error(w, "Error registering user", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, err = w.Write([]byte(`{"message":"User registered successfully"}`))
		if err != nil {
			http.Error(w, "Error writing response", http.StatusInternalServerError)
			return
		}
	}
}

func (a *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		user := domain.User{}
		err = json.Unmarshal(body, &user)
		if err != nil {
			http.Error(w, "Error parsing request body", http.StatusBadRequest)
			return
		}

		a.userService.Login(context.Background(), user)

		tokenString, err := a.tokenService.GenerateToken(user.Username)
		if err != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(`{"token":"` + tokenString + `"}`))
		if err != nil {
			http.Error(w, "Error writing response", http.StatusInternalServerError)
			return
		}
	}
}
