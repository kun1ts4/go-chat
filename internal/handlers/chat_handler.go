package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"go-chat/internal/domain"
	"go-chat/internal/services"
	"io"
	"net/http"
)

type ChatHandler struct {
	chatService *services.ChatService
}

func NewChatHandler(chatService *services.ChatService) *ChatHandler {
	return &ChatHandler{chatService: chatService}
}

func (c *ChatHandler) PostChatMessage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		message := domain.Message{}
		sender := r.Context().Value("user").(string)
		message.Sender = sender

		err = json.Unmarshal(body, &message)
		if err != nil {
			http.Error(w, "Error parsing request body", http.StatusBadRequest)
			return
		}

		err = c.chatService.PostChatMessage(context.Background(), message)
		if err != nil {
			http.Error(w, "failed to create message", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"message": "Message created successfully"`))
	}
}

func (c *ChatHandler) GetChat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		messages, err := c.chatService.GetChat(context.Background())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonResult, err := json.Marshal(messages)
		if err != nil {
			http.Error(w, "json marshaling error", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(fmt.Sprintf(`{"messages": %s}`, jsonResult)))
	}
}
