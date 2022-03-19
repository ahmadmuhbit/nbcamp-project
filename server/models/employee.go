package models

import (
	"github.com/google/uuid"
	"time"
)

type Employee struct {
	ID        uuid.UUID
	NIP       string
	Name      string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewEmployee() *Employee {
	return &Employee{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
