package handler

import (
	"net/http"
	"time"

	"github.com/davidchristie/app/services/server/auth"
	"github.com/davidchristie/app/services/server/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(config *config.Config, auth auth.Auth) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Mount("/api", apiHandler(auth))
	r.Get("/*", Web(config.WebDirectory))
	return r
}
