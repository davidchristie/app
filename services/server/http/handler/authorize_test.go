package handler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/davidchristie/app/services/server/http/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthorize(t *testing.T) {
	type req struct {
		providerID string
	}
	tests := []struct {
		name        string
		req         req
		wantHeaders map[string]string
		wantCode    int
		wantBody    string
	}{
		{
			name: "github",
			req: req{
				providerID: "github",
			},
			wantHeaders: map[string]string{
				"Location": "/",
			},
			wantCode: http.StatusMovedPermanently,
			wantBody: "<a href=\"/\">Moved Permanently</a>.\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			r, err := http.NewRequest("GET", fmt.Sprintf("/api/v1/auth/%s/authorize", tt.req.providerID), nil)
			require.NoError(t, err)
			handler.Authorize().ServeHTTP(rr, r)
			for key, value := range tt.wantHeaders {
				assert.Equal(t, value, rr.Header().Get(key))
			}
			assert.Equal(t, tt.wantCode, rr.Code)
			assert.Equal(t, tt.wantBody, rr.Body.String())
		})
	}
}
