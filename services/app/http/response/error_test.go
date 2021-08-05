package response_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/davidchristie/app/services/app/http/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody interface{}
	}{
		{
			name: "unknown_error",
			args: args{
				err: errors.New("unknown"),
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
			response.Error(rr, tt.args.err)
			assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
			assert.Equal(t, tt.wantCode, rr.Code)
			var body interface{}
			err := json.Unmarshal(rr.Body.Bytes(), &body)
			require.NoError(t, err)
			assert.Equal(t, tt.wantBody, body)
		})
	}
}
