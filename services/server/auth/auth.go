package auth

import (
	"context"

	"github.com/google/uuid"
)

type Auth interface {
	Session(ctx context.Context, token string) (*Session, error)
}

type Session struct {
	User *User
}

type User struct {
	ID    uuid.UUID
	Name  string
	Email string
}

type auth struct{}

func NewAuth() Auth {
	return &auth{}
}

func (a *auth) Session(ctx context.Context, token string) (*Session, error) {
	if token == "" {
		return &Session{
			User: nil,
		}, nil
	}
	return &Session{
		User: &User{
			ID:    uuid.MustParse("97406d59-7a49-4f1e-bb79-aba34cfcb405"),
			Name:  "Test User",
			Email: "test_user@email.com",
		},
	}, nil
}
