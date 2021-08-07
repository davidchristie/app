package handler_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/davidchristie/app/services/app/auth"
	"github.com/davidchristie/app/services/app/http/handler"
	"github.com/davidchristie/app/services/app/mocks"
	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAuthorize(t *testing.T) {
	type req struct {
		providerID string
	}
	ctrl := gomock.NewController(t)
	tests := []struct {
		name        string
		req         req
		auth        auth.Auth
		wantHeaders map[string]string
		wantCode    int
		wantBody    string
	}{
		{
			name: "github",
			req: req{
				providerID: "github",
			},
			auth: (func() auth.Auth {
				mock := mocks.NewMockAuth(ctrl)
				mock.EXPECT().Authorize("github").Return(&auth.AuthorizeResult{Redirect: "https://github.com/login/oauth/authorize"}, nil)
				return mock
			})(),
			wantHeaders: map[string]string{
				"Location": "https://github.com/login/oauth/authorize",
			},
			wantCode: http.StatusTemporaryRedirect,
			wantBody: "<a href=\"https://github.com/login/oauth/authorize\">Temporary Redirect</a>.\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			ctx := chi.NewRouteContext()
			ctx.URLParams.Add("providerID", tt.req.providerID)
			r := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/auth/%s/authorize", tt.req.providerID), nil)
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
			handler.Authorize(tt.auth).ServeHTTP(rr, r)
			for key, value := range tt.wantHeaders {
				assert.Equal(t, value, rr.Header().Get(key))
			}
			assert.Equal(t, tt.wantCode, rr.Code)
			assert.Equal(t, tt.wantBody, rr.Body.String())
		})
	}
}
