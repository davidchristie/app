package auth_test

import (
	"testing"

	"github.com/davidchristie/app/services/app/auth"
	"github.com/davidchristie/app/services/app/config"
	"github.com/davidchristie/app/services/app/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const (
	sessionToken = "dc33f5d8a6af86ad8964b62dc5ae6567"
	userName     = "Test User"
	userEmail    = "test_user@email.com"
)

var userID = uuid.MustParse("97406d59-7a49-4f1e-bb79-aba34cfcb405")

func TestNewAuth(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "default",
		},
	}
	ctrl := gomock.NewController(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auth := auth.NewAuth(
				config.DefaultConfig(),
				mocks.NewMockUserRepository(ctrl),
				mocks.NewMockAccountRepository(ctrl),
				mocks.NewMockSessionRepository(ctrl),
			)
			assert.NotNil(t, auth)
		})
	}
}
