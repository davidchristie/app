package handler

import (
	"net/http"

	"github.com/davidchristie/app/services/server/http/middleware"
)

func SignOut() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:     middleware.SessionTokenCookie,
			Value:    "",
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: true,
		})
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}
