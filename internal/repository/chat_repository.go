package repository

import (
	"context"
	"github.com/jackc/pgx/v4"
	"go-chat/internal/domain"
)

type ChatRepository interface {
	PostChatMessage(ctx context.Context, message domain.Message) error
	GetChat(ctx context.Context) ([]domain.Message, error)
}

type ChatRepo struct {
	db *pgx.Conn
}

func NewChatRepo(db *pgx.Conn) *ChatRepo {
	return &ChatRepo{db: db}
}

func (r *ChatRepo) PostChatMessage(ctx context.Context, message domain.Message) error {
	_, err := r.db.Exec(ctx, `INSERT INTO global_messages (sender, message_text) VALUES ($1, $2)`, message.Sender, message.Text)
	if err != nil {
		return err
	}
	return nil
}

func (r *ChatRepo) GetChat(ctx context.Context) ([]domain.Message, error) {
	q, err := r.db.Query(ctx, `SELECT (sender, message_text, created_at) FROM global_messages`)
	if err != nil {
		return nil, err
	}

	var messages []domain.Message
	for q.Next() {
		m := domain.Message{}
		err = q.Scan(&m.Sender, &m.Text, &m.CreatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, m)
	}
	return messages, nil
}
