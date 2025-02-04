package services

import (
	"context"
	"errors"
	"go-chat/internal/domain"
	"go-chat/internal/repository"
)

type ChatService struct {
	repo repository.ChatRepository
}

func NewChatService(repo repository.ChatRepository) *ChatService {
	return &ChatService{repo: repo}
}

func (c *ChatService) PostChatMessage(ctx context.Context, message domain.Message) error {
	if message.Text == "" {
		return errors.New("no message text provided")
	}
	if message.Sender == "" {
		return errors.New("username cant be empty")
	}
	return c.repo.PostChatMessage(ctx, message)
}

func (c *ChatService) GetChat(ctx context.Context) ([]domain.Message, error) {
	return c.repo.GetChat(ctx)
}
