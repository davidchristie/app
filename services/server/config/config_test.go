package config_test

import (
	"os"
	"testing"

	"github.com/davidchristie/app/services/server/config"
	"github.com/stretchr/testify/assert"
)

func TestDefaultConfig(t *testing.T) {
	assert.Equal(t, &config.Config{
		Port:         4000,
		WebDirectory: "./public",
	}, config.DefaultConfig())
}

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name    string
		env     map[string]string
		want    *config.Config
		wantErr string
	}{
		{
			name: "default",
			want: config.DefaultConfig(),
		},
		{
			name: "port_environment_variable",
			env: map[string]string{
				"PORT": "8080",
			},
			want: &config.Config{
				Port:         8080,
				WebDirectory: "./public",
			},
		},
		{
			name: "invalid_port",
			env: map[string]string{
				"PORT": "not_a_number",
			},
			want:    nil,
			wantErr: "envconfig.Process: assigning PORT to Port: converting 'not_a_number' to type int. details: strconv.ParseInt: parsing \"not_a_number\": invalid syntax",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Clearenv()
			for key, value := range tt.env {
				os.Setenv(key, value)
			}
			config, err := config.LoadConfig()
			assert.Equal(t, tt.want, config)
			if err != nil || tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
			}
		})
	}
}
