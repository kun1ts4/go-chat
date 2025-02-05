package domain

type DirectMessage struct {
	Recipient string `json:"recipient"`
	Message
}
