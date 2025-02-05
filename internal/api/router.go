package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go-chat/internal/handlers"
	mw "go-chat/internal/middleware"
	"go-chat/internal/services"
)

func NewRouter(authHandler *handlers.AuthHandler, tokenService *services.TokenService, chatHandler *handlers.ChatHandler, dmHandler *handlers.DMHandler) *chi.Mux {
	root := chi.NewRouter()
	root.Use(middleware.Logger)

	root.Post("/reg", authHandler.Register())
	root.Post("/login", authHandler.Login())

	r := chi.NewRouter()
	r.Use(mw.Auth(tokenService))

	r.Post("/chat", chatHandler.PostChatMessage())
	r.Get("/chat", chatHandler.GetChat())
	r.Get("/dms", dmHandler.GetUserDMs())
	r.Post("/dms", dmHandler.SendDirectMessage())

	root.Mount("/api", r)

	return root
}
