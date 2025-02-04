package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go-chat/internal/handlers"
	mw "go-chat/internal/middleware"
	"go-chat/internal/services"
	"net/http"
)

func NewRouter(authHandler *handlers.AuthHandler, tokenService *services.TokenService, chatHandler *handlers.ChatHandler) *chi.Mux {
	root := chi.NewRouter()
	root.Use(middleware.Logger)

	root.Post("/reg", authHandler.Register())
	root.Post("/login", authHandler.Login())

	r := chi.NewRouter()
	r.Use(mw.Auth(tokenService))
	r.Post("/chat", chatHandler.PostChatMessage())

	r.Get("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello"))
	})

	root.Mount("/api", r)

	return root
}
