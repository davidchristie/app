package config_test

import (
	"os"
	"testing"

	"github.com/davidchristie/app/services/app/config"
	"github.com/stretchr/testify/assert"
)

func TestDefaultConfig(t *testing.T) {
	assert.Equal(t, &config.Config{
		DatabaseMigrations: "file://services/app/migrations",
		DatabaseType:       "postgres",
		DatabaseURL:        "postgres://user:password@localhost:5432/db?sslmode=disable",
		GitHubRedirectURL:  "https://app-production-321806-o77vfhyfuq-ts.a.run.app/api/v1/auth/github/callback",
		GitHubAuthURL:      "https://github.com/login/oauth/authorize",
		GitHubTokenURL:     "https://github.com/login/oauth/access_token",
		GitHubUserURL:      "https://api.github.com/user",
		GitHubEmailsURL:    "https://api.github.com/user/emails",
		GoogleRedirectURL:  "https://app-production-321806-o77vfhyfuq-ts.a.run.app/api/v1/auth/google/callback",
		GoogleAuthURL:      "https://accounts.google.com/o/oauth2/auth",
		GoogleTokenURL:     "https://oauth2.googleapis.com/token",
		GoogleUserURL:      "https://www.googleapis.com/oauth2/v3/userinfo",
		Port:               4000,
		WebDirectory:       "./public",
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
			want: (func() *config.Config {
				config := config.DefaultConfig()
				config.Port = 8080
				return config
			})(),
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
