package handler

import (
	"net/http"

	"github.com/davidchristie/app/services/app/http/middleware"
	"github.com/davidchristie/app/services/app/http/response"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type sessionResponseBody struct {
	User *user `copier:"must,nopanic" json:"user"`
}

type user struct {
	ID    uuid.UUID `copier:"must,nopanic" json:"id"`
	Name  string    `copier:"must,nopanic" json:"name"`
	Email string    `copier:"must,nopanic" json:"email"`
}

func Session() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := middleware.GetSession(r.Context())
		if err != nil {
			response.Error(w, err)
			return
		}
		body := &sessionResponseBody{}
		if err = copier.Copy(body, session); err != nil {
			response.Error(w, err)
			return
		}
		response.JSON(w, body)
	}
}
