package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/davidchristie/app/services/server/auth"
	"github.com/davidchristie/app/services/server/http/response"
)

type contextKey string

const (
	SessionTokenCookie = "session-token"
	sessionContextKey  = contextKey("session")
)

var ErrNoSessionInContext = errors.New("no session in context")

func Session(auth auth.Auth) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			token := sessionToken(r)
			session, err := auth.Session(ctx, token)
			if err != nil {
				response.Error(w, err)
				return
			}
			next.ServeHTTP(w, r.WithContext(WithSession(ctx, session)))
		})
	}
}

func GetSession(ctx context.Context) (*auth.Session, error) {
	session, ok := ctx.Value(sessionContextKey).(*auth.Session)
	if !ok {
		return nil, ErrNoSessionInContext
	}
	return session, nil
}

func WithSession(ctx context.Context, session *auth.Session) context.Context {
	return context.WithValue(ctx, sessionContextKey, session)
}

func sessionToken(r *http.Request) string {
	cookie, _ := r.Cookie(SessionTokenCookie)
	if cookie != nil {
		return cookie.String()
	}
	return ""
}
