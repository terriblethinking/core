package store

import (
	"time"

	"github.com/maximhq/bifrost/core/schemas"
	"github.com/google/uuid"
)

type Chat struct {
	ID uuid.UUID
	Time time.Time
	Messages []schemas.ChatMessage
}

func NewChat(prompt string, history []schemas.ChatMessage) *Chat {

	ID := uuid.New()
	Time := time.Now()
	var Messages []schemas.ChatMessage

	if len(history) != 0 {
		Messages = history
	}

	return &Chat{
		ID: ID,
		Time: Time,
		Messages: Messages,
	}

}

