package main

import (
	"context"
	"go-chat/internal/api"
	"go-chat/internal/handlers"
	"go-chat/internal/repository"
	"go-chat/internal/services"
	"go-chat/pkg/postgres"
	"log"
	"net/http"
)

func main() {
	db, err := postgres.NewPostgresDatabase()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	defer db.Close(context.Background())

	userRepo := repository.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	tokenService := services.NewTokenService("secret")
	authHandler := handlers.NewAuthHandler(userService, tokenService)

	chatRepo := repository.NewChatRepo(db)
	chatService := services.NewChatService(chatRepo)
	chatHandler := handlers.NewChatHandler(chatService)

	dmRepo := repository.NewDMRepo(db)
	dmService := services.NewDMService(dmRepo)
	dmHandler := handlers.NewDMHandler(dmService)

	r := api.NewRouter(authHandler, tokenService, chatHandler, dmHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
