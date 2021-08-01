package handler_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/davidchristie/app/services/server/auth"
	"github.com/davidchristie/app/services/server/config"
	"github.com/davidchristie/app/services/server/http/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	htmlFilename = "index.html"
	htmlContent  = "<html><!-- ... --></html>"
)

func TestNewHandler(t *testing.T) {
	webDirectory, err := ioutil.TempDir("", "test_web_directory")
	require.NoError(t, err)
	defer os.RemoveAll(webDirectory)
	err = ioutil.WriteFile(filepath.Join(webDirectory, htmlFilename), []byte(htmlContent), 0666)
	require.NoError(t, err)
	config := config.DefaultConfig()
	config.WebDirectory = webDirectory
	type args struct {
		auth auth.Auth
	}
	type req struct {
		method string
		url    string
		body   map[string]interface{}
	}
	tests := []struct {
		name        string
		args        args
		req         req
		wantHeaders map[string]string
		wantCode    int
		wantBody    interface{}
	}{
		{
			name: "get_api_v1_auth_session",
			args: args{
				auth: auth.NewAuth(),
			},
			req: req{
				method: "GET",
				url:    "/api/v1/auth/session",
			},
			wantHeaders: map[string]string{
				"Content-Type": "application/json",
			},
			wantCode: http.StatusOK,
			wantBody: map[string]interface{}{
				"user": nil,
			},
		},
		{
			name: "get_api_unknown_url",
			args: args{
				auth: auth.NewAuth(),
			},
			req: req{
				method: "GET",
				url:    "/api/unknown/url",
			},
			wantCode: http.StatusNotFound,
			wantBody: "404 page not found\n",
		},
		{
			name: "get_web_html",
			args: args{
				auth: auth.NewAuth(),
			},
			req: req{
				method: "GET",
				url:    "/",
			},
			wantCode: http.StatusOK,
			wantBody: htmlContent,
		},
		{
			name: "get_web_root",
			args: args{
				auth: auth.NewAuth(),
			},
			req: req{
				method: "GET",
				url:    "/",
			},
			wantCode: http.StatusOK,
			wantBody: htmlContent,
		},
		{
			name: "get_web_index_html",
			args: args{
				auth: auth.NewAuth(),
			},
			req: req{
				method: "GET",
				url:    "/index.html",
			},
			wantHeaders: map[string]string{
				"Location": "./",
			},
			wantCode: http.StatusMovedPermanently,
			wantBody: "",
		},
		{
			name: "get_web_other_url",
			args: args{
				auth: auth.NewAuth(),
			},
			req: req{
				method: "GET",
				url:    "/other/url",
			},
			wantCode: http.StatusOK,
			wantBody: htmlContent,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := handler.NewHandler(config, tt.args.auth)
			rr := httptest.NewRecorder()
			reqBody, err := json.Marshal(tt.req.body)
			require.NoError(t, err)
			req, err := http.NewRequest(tt.req.method, tt.req.url, bytes.NewBuffer(reqBody))
			require.NoError(t, err)
			handler.ServeHTTP(rr, req)
			for key, value := range tt.wantHeaders {
				assert.Equal(t, value, rr.Header().Get(key))
			}
			assert.Equal(t, tt.wantCode, rr.Code)
			var body interface{} = rr.Body.String()
			if rr.Header().Get("Content-Type") == "application/json" {
				err = json.Unmarshal(rr.Body.Bytes(), &body)
				require.NoError(t, err)
			}
			assert.Equal(t, tt.wantBody, body)
		})
	}
}
