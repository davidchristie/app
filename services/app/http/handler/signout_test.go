package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/davidchristie/app/services/app/http/handler"
	"github.com/stretchr/testify/assert"
)

func TestSignOut(t *testing.T) {
	tests := []struct {
		name        string
		wantHeaders map[string]string
		wantCode    int
		wantBody    string
	}{
		{
			name: "default",
			wantHeaders: map[string]string{
				"Location":   "/",
				"Set-Cookie": "session-token=; Path=/; Max-Age=0; HttpOnly",
			},
			wantCode: http.StatusMovedPermanently,
			wantBody: "<a href=\"/\">Moved Permanently</a>.\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/api/v1/auth/signout", nil)
			handler.SignOut().ServeHTTP(rr, r)
			for key, value := range tt.wantHeaders {
				assert.Equal(t, value, rr.Header().Get(key))
			}
			assert.Equal(t, tt.wantCode, rr.Code)
			assert.Equal(t, tt.wantBody, rr.Body.String())
		})
	}
}
