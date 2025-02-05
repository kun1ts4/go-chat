package domain

import "time"

type Message struct {
	Text      string    `json:"text"`
	Sender    string    `json:"sender"`
	CreatedAt time.Time `json:"createdAt"`
}
