package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/davidchristie/app/services/server/auth"
	"github.com/davidchristie/app/services/server/config"
	"github.com/davidchristie/app/services/server/http/handler"
)

type Server interface {
	Close() error
	Start() error
}

type server struct {
	base *http.Server
}

func NewServer(config *config.Config, auth auth.Auth) Server {
	return &server{
		base: &http.Server{
			Addr:    fmt.Sprintf(":%d", config.Port),
			Handler: handler.NewHandler(config, auth),
		},
	}
}

func (s *server) Close() error {
	log.Println("Closing server")
	return s.base.Close()
}

func (s *server) Start() error {
	log.Println("Starting server on", s.base.Addr)
	return s.base.ListenAndServe()
}
