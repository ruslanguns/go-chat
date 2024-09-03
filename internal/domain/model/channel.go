package model

import "github.com/ruslanguns/go-chat/internal/domain"

type Channel struct {
	domain.BaseEntity
	Name        string `gorm:"uniqueIndex" json:"name"`
	Description string `json:"description"`
	Users       []User `gorm:"many2many:user_channels;" json:"users"`
}
