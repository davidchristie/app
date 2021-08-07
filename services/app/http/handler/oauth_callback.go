package handler

import (
	"net/http"

	"github.com/davidchristie/app/services/app/auth"
	"github.com/davidchristie/app/services/app/http/middleware"
	"github.com/davidchristie/app/services/app/http/response"
	"github.com/go-chi/chi/v5"
)

func OAuthCallback(auth auth.Auth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		providerID := chi.URLParam(r, "providerID")
		code := r.URL.Query().Get("code")
		state := r.URL.Query().Get("state")
		result, err := auth.Callback(r.Context(), providerID, state, code)
		if err != nil {
			response.Error(w, err)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     middleware.SessionTokenCookie,
			Value:    result.SessionToken,
			Path:     "/",
			Expires:  result.SessionExpiresAt,
			HttpOnly: true,
		})
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}
