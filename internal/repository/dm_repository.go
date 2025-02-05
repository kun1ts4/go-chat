package repository

import (
	"context"
	"github.com/jackc/pgx/v4"
	"go-chat/internal/domain"
)

type DMRepository interface {
	SendDirectMessage(ctx context.Context, message domain.DirectMessage) error
	GetUserDMs(ctx context.Context, username string) ([]domain.DirectMessage, error)
}

type DMRepo struct {
	db *pgx.Conn
}

func NewDMRepo(db *pgx.Conn) *DMRepo {
	return &DMRepo{db: db}
}

func (r *DMRepo) SendDirectMessage(ctx context.Context, message domain.DirectMessage) error {
	_, err := r.db.Exec(ctx, `INSERT INTO direct_messages (sender, receiver, message_text) VALUES ($1, $2, $3)`, message.Sender, message.Recipient, message.Text)
	if err != nil {
		return err
	}

	return nil
}

func (r *DMRepo) GetUserDMs(ctx context.Context, username string) ([]domain.DirectMessage, error) {
	q, err := r.db.Query(ctx, `SELECT sender, receiver, message_text, created_at FROM direct_messages WHERE receiver = $1`, username)
	if err != nil {
		return nil, err
	}
	var userDMs []domain.DirectMessage

	for q.Next() {
		message := domain.DirectMessage{}
		err := q.Scan(&message.Sender, &message.Recipient, &message.Text, &message.CreatedAt)
		if err != nil {
			return nil, err
		}
		userDMs = append(userDMs, message)
	}

	return userDMs, nil
}
