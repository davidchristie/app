package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	PrimaryEmail string
	FullName     string
	AvatarURL    string
}
