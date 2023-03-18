package models

type Message struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Length  uint64 `json:"length"`
}
