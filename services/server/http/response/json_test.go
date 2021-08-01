package response_test

import (
	"encoding/json"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/davidchristie/app/services/server/http/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	validJSON = map[string]interface{}{
		"boolean": true,
		"number":  123.0,
		"string":  "abc",
		"object": map[string]interface{}{
			"key": "value",
		},
	}
)

func TestJSON(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody interface{}
	}{
		{
			name: "valid_json",
			args: args{
				v: validJSON,
			},
			wantCode: http.StatusOK,
			wantBody: validJSON,
		},
		{
			name: "unsupported_type",
			args: args{
				v: func() {},
			},
			wantCode: http.StatusInternalServerError,
			wantBody: map[string]interface{}{
				"error": "Something went wrong",
			},
		},
		{
			name: "unsupported_value",
			args: args{
				v: math.Inf(1),
			},
			wantCode: http.StatusInternalServerError,
			wantBody: map[string]interface{}{
				"error": "Something went wrong",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			response.JSON(rr, tt.args.v)
			assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
			assert.Equal(t, tt.wantCode, rr.Code)
			var body interface{}
			err := json.Unmarshal(rr.Body.Bytes(), &body)
			require.NoError(t, err)
			assert.Equal(t, tt.wantBody, body)
		})
	}
}
