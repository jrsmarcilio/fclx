package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	tiktoken_go "github.com/j178/tiktoken-go"
)

type Message struct {
	Content   string
	CreatedAt time.Time
	ID        string
	Model     *Model
	Role      string
	Tokens    int
}

func NewMessage(role, content string, model *Model) (*Message, error) {
	totalTokens := tiktoken_go.CountTokens(model.GetModelName(), content)
	msg := &Message{
		Content:   content,
		CreatedAt: time.Now(),
		ID:        uuid.New().String(),
		Model:     model,
		Role:      role,
		Tokens:    totalTokens,
	}
	if err := msg.Validate(); err != nil {
		return nil, err
	}
	return msg, nil
}

func (m *Message) Validate() error {
	if m.Role != "user" && m.Role != "system" && m.Role != "assistant" {
		return errors.New("invalid role")
	}
	if m.Content == "" {
		return errors.New("content is required")
	}
	if m.CreatedAt.IsZero() {
		return errors.New("created_at is required")
	}
	return nil
}
