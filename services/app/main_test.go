package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type MainTestSuite struct {
	suite.Suite
}

var (
	ErrServerNotCreated = errors.New("server not created")
	exitCode            = 0
)

func Test_main(t *testing.T) {
	mockExit()
	port := freePort(t)
	env := map[string]string{
		"PORT": fmt.Sprint(port),
	}
	err := runServer(t, env, func() {
		res, err := http.Get(fmt.Sprintf("http://localhost:%d/api/v1/auth/session", port))
		require.NoError(t, err)
		defer res.Body.Close()
		var body interface{}
		err = json.NewDecoder(res.Body).Decode(&body)
		require.NoError(t, err)
		assert.Equal(t, map[string]interface{}{
			"user": nil,
		}, body)
	})
	require.NoError(t, err)
}

func freePort(t *testing.T) int {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	require.NoError(t, err)
	listener, err := net.ListenTCP("tcp", addr)
	require.NoError(t, err)
	defer listener.Close()
	return listener.Addr().(*net.TCPAddr).Port
}

func mockExit() {
	exit = func(code int) {
		exitCode = code
	}
}

func runServer(t *testing.T, env map[string]string, fn func()) error {
	setEnv(t, env)
	initServer()
	if server == nil {
		return ErrServerNotCreated
	}
	serverRunning := make(chan struct{})
	serverDone := make(chan struct{})
	go func() {
		close(serverRunning)
		main()
		defer close(serverDone)
	}()
	<-serverRunning
	fn()
	server.Close()
	<-serverDone
	assert.Equal(t, 1, exitCode)
	return nil
}

func setEnv(t *testing.T, env map[string]string) {
	os.Clearenv()
	for key, value := range env {
		err := os.Setenv(key, value)
		require.NoError(t, err)
	}
}
