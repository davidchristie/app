package handler

import (
	"net/http"

	"github.com/davidchristie/app/services/app/auth"
	"github.com/davidchristie/app/services/app/http/response"
	"github.com/go-chi/chi/v5"
)

func Authorize(auth auth.Auth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		providerID := chi.URLParam(r, "providerID")
		result, err := auth.Authorize(providerID)
		if err != nil {
			response.Error(w, err)
			return
		}
		http.Redirect(w, r, result.Redirect, http.StatusTemporaryRedirect)
	}
}
