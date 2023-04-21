package entity

import "errors"

type ChatConfig struct {
	Model            *Model
	Temperature      float32
	TopP             float32
	N                int
	Stop             []string
	MaxTokens        int
	PresencePenalty  float32
	FrequencyPenalty float32
}

type Chat struct {
	ID                   string
	UserID               string
	InitialSystemMessage *Message
	Messages             []*Message
	ErasedMessages       []*Message
	Status               string
	TokenUsage           int
	Config               *ChatConfig
}

func NewChat(userID string, initialSystemMessage *Message, config *ChatConfig) (*Chat, error) {
	chat := &Chat{
		ID:                   userID,
		UserID:               userID,
		InitialSystemMessage: initialSystemMessage,
		Status:               "active",
		Config:               config,
		TokenUsage:           0,
	}
	chat.AddMessage(initialSystemMessage)

	if err := chat.Validate(); err != nil {
		return nil, err
	}
	return chat, nil
}

func (chat *Chat) AddMessage(msg *Message) error {
	if chat.Status == "ended" {
		return errors.New("chat is ended. no more messages can be added")
	}

	for {
		if chat.Config.Model.GetMaxTokens() >= msg.GetQtdTokens()+chat.TokenUsage {
			chat.Messages = append(chat.Messages, msg)
			chat.RefreshTokenUsage()
			break
		}
		chat.ErasedMessages = append(chat.ErasedMessages, chat.Messages[0])
		chat.Messages = chat.Messages[1:]
		chat.RefreshTokenUsage()
	}
	return nil
}

func (chat *Chat) GetMessages() []*Message {
	return chat.Messages
}

func (chat *Chat) CountMessages() int {
	return len(chat.Messages)
}

func (chat *Chat) End() {
	chat.Status = "ended"
}

func (chat *Chat) RefreshTokenUsage() {
	chat.TokenUsage = 0
	for msg := range chat.Messages {
		chat.TokenUsage += chat.Messages[msg].GetQtdTokens()
	}
}

func (chat *Chat) Validate() error {
	if chat.UserID == "" {
		return errors.New("user_id is required")
	}
	if chat.Status != "active" && chat.Status != "ended" {
		return errors.New("invalid status")
	}
	if chat.Config.Temperature < 0 || chat.Config.Temperature > 2 {
		return errors.New("invalid temperature")
	}
	// more validations for config
	if chat.InitialSystemMessage == nil {
		return errors.New("initial_system_message is required")
	}
	if chat.Config == nil {
		return errors.New("config is required")
	}
	return nil
}
