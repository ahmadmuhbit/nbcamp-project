package models

import (
	"github.com/google/uuid"
	"time"
)

type Menu struct {
	ID          uuid.UUID
	Name        string
	Category    string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewMenu() *Menu {
	return &Menu{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
