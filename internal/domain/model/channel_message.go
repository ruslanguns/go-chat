package model

import (
	"github.com/ruslanguns/go-chat/internal/domain"
)

type ChannelMessage struct {
	domain.BaseEntity
	ChannelID domain.EntityID `json:"channel_id"`
	SenderID  domain.EntityID `json:"sender_id"`
	Content   string          `json:"content"`
}
