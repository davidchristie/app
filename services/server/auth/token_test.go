package auth

import (
	"bytes"
	"compress/flate"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateSecureToken(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr string
	}{
		{
			name: "length_32",
			args: args{
				length: 32,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := GenerateSecureToken(tt.args.length)
			if err == nil {
				testCompression(t, token)
			}
			if err != nil || tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
			}
		})
	}
}

func testCompression(t *testing.T, token string) {
	// Based on the tests for the rand package.
	// https://cs.opensource.google/go/go/+/refs/tags/go1.16.6:src/crypto/rand/rand_test.go
	t.Run("compression", func(t *testing.T) {
		b, err := hex.DecodeString(token)
		if err != nil {
			require.NoError(t, err)
		}
		var z bytes.Buffer
		f, _ := flate.NewWriter(&z, 5)
		f.Write(b)
		f.Close()
		assert.Greater(t, z.Len(), len(b)*99/100)
	})
}
