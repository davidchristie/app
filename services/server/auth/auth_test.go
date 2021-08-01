package auth_test

import (
	"context"
	"testing"

	"github.com/davidchristie/app/services/server/auth"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const (
	sessionToken = "dc33f5d8a6af86ad8964b62dc5ae6567"
	userName     = "Test User"
	userEmail    = "test_user@email.com"
)

var userID = uuid.MustParse("97406d59-7a49-4f1e-bb79-aba34cfcb405")

func TestNewAuth(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "default",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auth := auth.NewAuth()
			assert.NotNil(t, auth)
			testSession(t, auth)
		})
	}
}

func testSession(t *testing.T, a auth.Auth) {
	t.Run("Session", func(t *testing.T) {
		type args struct {
			ctx   context.Context
			token string
		}
		tests := []struct {
			name    string
			args    args
			want    *auth.Session
			wantErr string
		}{
			{
				name: "valid_token",
				args: args{
					ctx:   context.Background(),
					token: sessionToken,
				},
				want: &auth.Session{
					User: &auth.User{
						ID:    userID,
						Name:  userName,
						Email: userEmail,
					},
				},
			},
			{
				name: "no_token",
				args: args{
					ctx: context.Background(),
				},
				want: &auth.Session{
					User: nil,
				},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				session, err := a.Session(tt.args.ctx, tt.args.token)
				assert.Equal(t, tt.want, session)
				if err != nil || tt.wantErr != "" {
					assert.EqualError(t, err, tt.wantErr)
				}
			})
		}
	})
}
