package domain

import "time"

type Message struct {
	Text      string `json:"text"`
	Sender    string
	CreatedAt time.Time
}
