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

type DMHandler struct {
	dmService *services.DMService
}

func NewDMHandler(dmService *services.DMService) *DMHandler {
	return &DMHandler{dmService: dmService}
}

func (d *DMHandler) SendDirectMessage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		message := domain.DirectMessage{}
		message.Sender = r.Context().Value("user").(string)
		err = json.Unmarshal(body, &message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = d.dmService.SendDirectMessage(context.Background(), message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(fmt.Sprintf(`{"message": "ditect message sent"}`)))
	}
}

func (c *DMHandler) GetUserDMs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.Context().Value("user").(string)
		messages, err := c.dmService.GetUserDMs(context.Background(), username)
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
