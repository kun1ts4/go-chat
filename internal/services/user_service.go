package services

import (
	"context"
	"errors"
	"go-chat/internal/domain"
	"go-chat/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(ctx context.Context, user domain.User) error {
	existingUser, _ := s.repo.GetByUsername(ctx, user.Username)
	if existingUser.Username != "" {
		return errors.New("user already exists")
	}

	return s.repo.Create(ctx, user)
}

func (s *UserService) Login(ctx context.Context, user domain.User) (domain.User, error) {
	existingUser, err := s.repo.GetByUsername(ctx, user.Username)
	if err != nil {
		return domain.User{}, err
	}

	if existingUser.Username == "" {
		return domain.User{}, errors.New("user does not exist")
	}

	if existingUser.Password != user.Password {
		return domain.User{}, errors.New("invalid password")
	}

	return user, nil
}
