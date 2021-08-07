package entities

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID                uuid.UUID
	CreatedAt         time.Time
	UpdatedAt         time.Time
	ProviderType      string
	ProviderID        string
	ProviderAccountID string
	UserID            uuid.UUID
}
