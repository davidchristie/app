package handler

import (
	"net/http"

	"github.com/davidchristie/app/services/app/http/middleware"
	"github.com/go-chi/chi/v5"
)

func Authorize() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		providerID := chi.URLParam(r, "providerID")
		token := "f225cab4aa518b34f6dd24fdc665c338a43c979c50d24b3a4ae7eb078cd7cbbb"
		http.SetCookie(w, &http.Cookie{
			Name:     middleware.SessionTokenCookie,
			Value:    providerID + "_" + token,
			Path:     "/",
			MaxAge:   7 * 24 * 60 * 60,
			HttpOnly: true,
		})
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}
