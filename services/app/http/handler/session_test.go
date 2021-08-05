package handler_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/davidchristie/app/services/app/auth"
	"github.com/davidchristie/app/services/app/http/handler"
	"github.com/davidchristie/app/services/app/http/middleware"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	userID                            = uuid.MustParse("48d3f72a-c137-4b37-85ab-f3aca4994855")
	userName                          = "Test User"
	userEmail                         = "test_user@email.com"
	userWithMissingFields interface{} = map[string]interface{}{
		"ID": userID,
	}
)

func TestSession(t *testing.T) {
	tests := []struct {
		name     string
		ctx      context.Context
		wantCode int
		wantBody interface{}
	}{
		{
			name: "authenticated",
			ctx: middleware.WithSession(context.Background(), &auth.Session{
				User: &auth.User{
					ID:    userID,
					Name:  userName,
					Email: userEmail,
				},
			}),
			wantCode: http.StatusOK,
			wantBody: map[string]interface{}{
				"user": map[string]interface{}{
					"id":    userID.String(),
					"name":  userName,
					"email": userEmail,
				},
			},
		},
		{
			name: "unauthenticated",
			ctx: middleware.WithSession(context.Background(), &auth.Session{
				User: nil,
			}),
			wantCode: http.StatusOK,
			wantBody: map[string]interface{}{
				"user": nil,
			},
		},
		{
			name:     "no_session_in_context",
			ctx:      context.Background(),
			wantCode: http.StatusInternalServerError,
			wantBody: map[string]interface{}{
				"error": "Something went wrong",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			session := handler.Session()
			handler := http.HandlerFunc(session)
			req, err := http.NewRequestWithContext(tt.ctx, "GET", "/api/v1/auth/session", nil)
			require.NoError(t, err)
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)
			assert.Equal(t, tt.wantCode, rr.Code)
			var body interface{}
			err = json.Unmarshal(rr.Body.Bytes(), &body)
			require.NoError(t, err)
			assert.Equal(t, tt.wantBody, body)
		})
	}
}
