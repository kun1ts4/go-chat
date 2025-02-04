package postgres

import (
	"context"
	"github.com/jackc/pgx/v4"
	"log"
)

func NewPostgresDatabase() (*pgx.Conn, error) {
	dsn := "postgres://user:pass@localhost:5452/chat?sslmode=disable"
	db, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
		return nil, err
	}

	return db, nil
}
