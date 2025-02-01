package main

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jackc/pgx/v4"
	"io"
	"log"
	"net/http"
)

func main() {
	dsn := "postgres://user:pass@localhost:5452/chat?sslmode=disable"
	db, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	defer db.Close(context.Background())

	root := chi.NewRouter()
	root.Use(middleware.Logger)

	root.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message":"Hello world"}`))
	})
	root.Post("/reg", Register(db))

	log.Fatal(http.ListenAndServe(":8080", root))
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		user := User{}
		err = json.Unmarshal(body, &user)
		if err != nil {
			http.Error(w, "Error parsing request body", http.StatusBadRequest)
			return
		}

		exists := false
		err = db.QueryRow(context.Background(), `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)`, user.Username).Scan(&exists)
		if err != nil {
			http.Error(w, "Error checking user", http.StatusInternalServerError)
			return
		}

		if exists {
			http.Error(w, "User already exists", http.StatusConflict)
			return
		}

		_, err = db.Exec(context.Background(), `INSERT INTO users (username, password) VALUES ($1, $2)`, user.Username, user.Password)
		if err != nil {
			http.Error(w, "Error inserting user", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, err = w.Write([]byte(`{"message":"User registered successfully"}`))
		if err != nil {
			http.Error(w, "Error writing response", http.StatusInternalServerError)
			return
		}
	}
}
