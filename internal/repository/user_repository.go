package repository

import (
	"context"
	"github.com/jackc/pgx/v4"
	"go-chat/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user domain.User) error
	GetByUsername(ctx context.Context, username string) (domain.User, error)
}

type UserRepo struct {
	db *pgx.Conn
}

func NewUserRepo(db *pgx.Conn) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(ctx context.Context, user domain.User) error {
	_, err := r.db.Exec(ctx, "INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, user.Password)
	return err
}

func (r *UserRepo) GetByUsername(ctx context.Context, username string) (domain.User, error) {
	var user domain.User
	err := r.db.QueryRow(ctx, "SELECT * FROM users WHERE username = $1", username).Scan(&user.Username, &user.Password)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
