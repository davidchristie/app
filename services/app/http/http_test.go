package http_test

import (
	"testing"

	"github.com/davidchristie/app/services/app/auth"
	"github.com/davidchristie/app/services/app/config"
	"github.com/davidchristie/app/services/app/http"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	server := http.NewServer(config.DefaultConfig(), auth.NewAuth())
	assert.NotNil(t, server)
}
