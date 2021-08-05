package repositories

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/davidchristie/app/services/app/auth"
	"github.com/davidchristie/app/services/app/config"
	"github.com/davidchristie/app/services/app/database"
	"github.com/davidchristie/app/services/app/repositories"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func assertAccountEqual(t *testing.T, expected, actual *repositories.Account) {
	assert.Equal(t, expected.ID.String(), actual.ID.String())
	assert.Equal(t, expected.CreatedAt.Round(time.Microsecond).String(), actual.CreatedAt.String())
	assert.Equal(t, expected.UpdatedAt.Round(time.Microsecond).String(), actual.UpdatedAt.String())
	assert.Equal(t, expected.ProviderType, actual.ProviderType)
	assert.Equal(t, expected.ProviderID, actual.ProviderID)
	assert.Equal(t, expected.ProviderAccountID, actual.ProviderAccountID)
	assert.Equal(t, expected.UserID.String(), actual.UserID.String())
}

func assertSessionEqual(t *testing.T, expected, actual *repositories.Session) {
	assert.Equal(t, expected.ID.String(), actual.ID.String())
	assert.Equal(t, expected.CreatedAt.Round(time.Microsecond).String(), actual.CreatedAt.String())
	assert.Equal(t, expected.UpdatedAt.Round(time.Microsecond).String(), actual.UpdatedAt.String())
	assert.Equal(t, expected.ExpiresAt.Round(time.Microsecond).String(), actual.ExpiresAt.String())
	assert.Equal(t, expected.SessionToken, actual.SessionToken)
	assert.Equal(t, expected.UserID.String(), actual.UserID.String())
}

func assertUserEqual(t *testing.T, expected, actual *repositories.User) {
	assert.Equal(t, expected.ID.String(), actual.ID.String())
	assert.Equal(t, expected.CreatedAt.Round(time.Microsecond).String(), actual.CreatedAt.String())
	assert.Equal(t, expected.UpdatedAt.Round(time.Microsecond).String(), actual.UpdatedAt.String())
	assert.Equal(t, expected.PrimaryEmail, actual.PrimaryEmail)
	assert.Equal(t, expected.FullName, actual.FullName)
	assert.Equal(t, expected.AvatarURL, actual.AvatarURL)
}

func connectToDatabase(t *testing.T) *sql.DB {
	config := config.DefaultConfig()
	config.DatabaseMigrations = "file:../../../../services/app/migrations"
	db, err := database.NewConnection(config)
	require.NoError(t, err)
	return db
}

func randomAccount(userID uuid.UUID) *repositories.Account {
	return &repositories.Account{
		ID:                uuid.New(),
		CreatedAt:         time.Now().UTC(),
		UpdatedAt:         time.Now().UTC(),
		ProviderType:      "oauth2",
		ProviderID:        "github",
		ProviderAccountID: uuid.New().String(),
		UserID:            userID,
	}
}

func randomSession(t *testing.T, userID uuid.UUID) *repositories.Session {
	sessionToken, err := auth.GenerateSecureToken(32)
	require.NoError(t, err)
	return &repositories.Session{
		ID:           uuid.New(),
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
		ExpiresAt:    time.Now().Add(time.Hour * 24 * 7).UTC(),
		SessionToken: sessionToken,
		UserID:       userID,
	}
}

func randomUser() *repositories.User {
	return &repositories.User{
		ID:           uuid.New(),
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
		PrimaryEmail: fmt.Sprintf("test_user+%s@email.com", uuid.New().String()),
		FullName:     "Test User",
		AvatarURL:    "https://via.placeholder.com/150",
	}
}
