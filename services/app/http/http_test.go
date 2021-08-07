package http_test

import (
	"testing"

	"github.com/davidchristie/app/services/app/auth"
	"github.com/davidchristie/app/services/app/config"
	"github.com/davidchristie/app/services/app/http"
	"github.com/davidchristie/app/services/app/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	ctrl := gomock.NewController(t)
	server := http.NewServer(config.DefaultConfig(), auth.NewAuth(
		config.DefaultConfig(),
		mocks.NewMockUserRepository(ctrl),
		mocks.NewMockAccountRepository(ctrl),
		mocks.NewMockSessionRepository(ctrl),
	))
	assert.NotNil(t, server)
}
