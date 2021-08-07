package handler

import (
	"net/http"

	"github.com/davidchristie/app/services/app/auth"
	"github.com/davidchristie/app/services/app/http/middleware"
	"github.com/go-chi/chi/v5"
)

func apiHandler(auth auth.Auth) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Session(auth))
	r.Mount("/v1", restV1Handler(auth))
	return r
}

func restV1Handler(auth auth.Auth) http.Handler {
	r := chi.NewRouter()
	r.Get("/auth/session", Session())
	r.Get("/auth/{providerID}/authorize", Authorize(auth))
	r.Get("/auth/{providerID}/callback", OAuthCallback(auth))
	r.Get("/auth/signout", SignOut())
	return r
}
