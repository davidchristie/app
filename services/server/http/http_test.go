package http_test

import (
	"testing"

	"github.com/davidchristie/app/services/server/auth"
	"github.com/davidchristie/app/services/server/config"
	"github.com/davidchristie/app/services/server/http"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	server := http.NewServer(config.DefaultConfig(), auth.NewAuth())
	assert.NotNil(t, server)
}
