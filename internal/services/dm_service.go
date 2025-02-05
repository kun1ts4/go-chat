package services

import (
	"context"
	"go-chat/internal/domain"
	"go-chat/internal/repository"
)

type DMService struct {
	repo repository.DMRepository
}

func NewDMService(repo repository.DMRepository) *DMService {
	return &DMService{repo: repo}
}

func (d *DMService) SendDirectMessage(ctx context.Context, message domain.DirectMessage) error {
	return d.repo.SendDirectMessage(ctx, message)
}

func (d *DMService) GetUserDMs(ctx context.Context, username string) ([]domain.DirectMessage, error) {
	return d.repo.GetUserDMs(ctx, username)
}
