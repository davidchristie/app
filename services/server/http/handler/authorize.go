package handler

import (
	"net/http"
	"time"

	"github.com/davidchristie/app/services/server/auth"
	"github.com/davidchristie/app/services/server/http/middleware"
	"github.com/davidchristie/app/services/server/http/response"
	"github.com/go-chi/chi/v5"
)

func Authorize() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		providerID := chi.URLParam(r, "providerID")
		token, err := auth.GenerateSecureToken(32)
		if err != nil {
			response.Error(w, err)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     middleware.SessionTokenCookie,
			Value:    providerID + "_" + token,
			Path:     "/",
			Expires:  time.Now().Add(7 * 24 * time.Hour),
			HttpOnly: true,
		})
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}
