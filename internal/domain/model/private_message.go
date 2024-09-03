package model

import (
	"time"

	"github.com/ruslanguns/go-chat/internal/domain"
)

type PrivateMessage struct {
	domain.BaseEntity
	SenderID   domain.EntityID `json:"sender_id"`
	ReceiverID domain.EntityID `json:"receiver_id"`
	Content    string          `json:"content"`
	ReadAt     *time.Time      `json:"read_at"`
}
