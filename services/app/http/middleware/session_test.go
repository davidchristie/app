package middleware_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/davidchristie/app/services/app/auth"
	"github.com/davidchristie/app/services/app/http/middleware"
	"github.com/davidchristie/app/services/app/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	sessionToken = "f6a6d81bcfcde3f9d567db195b803ec6384f91e8802fc5d5ba8a8d5ff514b9e1"
)

var user = &auth.User{
	ID:    uuid.MustParse("97406d59-7a49-4f1e-bb79-aba34cfcb405"),
	Name:  "Test User",
	Email: "test_user@email.com",
}

func TestSession(t *testing.T) {
	type args struct {
		auth auth.Auth
	}
	type req struct {
		ctx   context.Context
		token string
	}
	ctrl := gomock.NewController(t)
	tests := []struct {
		name        string
		args        args
		req         req
		wantSession *auth.Session
	}{
		{
			name: "valid_session_token",
			args: args{
				auth: (func() auth.Auth {
					mock := mocks.NewMockAuth(ctrl)
					mock.EXPECT().Session(context.Background(), sessionToken).Return(&auth.Session{User: user}, nil)
					return mock
				})(),
			},
			req: req{
				ctx:   context.Background(),
				token: sessionToken,
			},
			wantSession: &auth.Session{
				User: user,
			},
		},
		{
			name: "no_session_token",
			args: args{
				auth: (func() auth.Auth {
					mock := mocks.NewMockAuth(ctrl)
					mock.EXPECT().Session(context.Background(), "").Return(&auth.Session{User: nil}, nil)
					return mock
				})(),
			},
			req: req{
				ctx: context.Background(),
			},
			wantSession: &auth.Session{
				User: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var nextReq *http.Request
			handler := middleware.Session(tt.args.auth)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				nextReq = r
			}))
			rr := httptest.NewRecorder()
			req, err := http.NewRequestWithContext(tt.req.ctx, "GET", "/", nil)
			if tt.req.token != "" {
				req.AddCookie(&http.Cookie{
					Name:  middleware.SessionTokenCookie,
					Value: tt.req.token,
				})
			}
			require.NoError(t, err)
			handler.ServeHTTP(rr, req)
			session, err := middleware.GetSession(nextReq.Context())
			require.NoError(t, err)
			assert.Equal(t, tt.wantSession, session)
		})
	}
}

func getSessionToken(r *http.Request) string {
	cookie, _ := r.Cookie(middleware.SessionTokenCookie)
	if cookie != nil {
		return cookie.String()
	}
	return ""
}
