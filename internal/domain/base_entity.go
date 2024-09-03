package domain

import (
	"time"

	"gorm.io/gorm"
)

type BaseEntity struct {
	ID        EntityID       `gorm:"type:string;primary_key" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (e *BaseEntity) BeforeCreate(tx *gorm.DB) error {
	if e.ID.IsZero() {
		e.ID = NewEntityID()
	}
	now := time.Now()
	e.CreatedAt = now
	e.UpdatedAt = now
	return nil
}

func (e *BaseEntity) BeforeUpdate(tx *gorm.DB) error {
	e.UpdatedAt = time.Now()
	return nil
}
