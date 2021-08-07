package entities

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID           uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ExpiresAt    time.Time
	SessionToken string
	UserID       uuid.UUID
}
