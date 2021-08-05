package handler_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/davidchristie/app/services/app/http/handler"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
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
				"Location":   "/",
				"Set-Cookie": "session-token=github_f225cab4aa518b34f6dd24fdc665c338a43c979c50d24b3a4ae7eb078cd7cbbb; Path=/; Max-Age=604800; HttpOnly",
			},
			wantCode: http.StatusMovedPermanently,
			wantBody: "<a href=\"/\">Moved Permanently</a>.\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			ctx := chi.NewRouteContext()
			ctx.URLParams.Add("providerID", tt.req.providerID)
			r := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/auth/%s/authorize", tt.req.providerID), nil)
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
			handler.Authorize().ServeHTTP(rr, r)
			for key, value := range tt.wantHeaders {
				assert.Equal(t, value, rr.Header().Get(key))
			}
			assert.Equal(t, tt.wantCode, rr.Code)
			assert.Equal(t, tt.wantBody, rr.Body.String())
		})
	}
}
