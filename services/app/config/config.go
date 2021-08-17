package config

import (
	"github.com/kelseyhightower/envconfig"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

type Config struct {
	DatabaseMigrations string `split_words:"true"`
	DatabaseType       string `split_words:"true"`
	DatabaseURL        string `split_words:"true"`
	GitHubClientID     string `envconfig:"github_client_id"`
	GitHubClientSecret string `envconfig:"github_client_secret"`
	GitHubRedirectURL  string `envconfig:"github_redirect_url"`
	GitHubAuthURL      string `envconfig:"github_auth_url"`
	GitHubTokenURL     string `envconfig:"github_token_url"`
	GitHubUserURL      string `envconfig:"github_user_url"`
	GitHubEmailsURL    string `envconfig:"github_emails_url"`
	GoogleClientID     string `split_words:"true"`
	GoogleClientSecret string `split_words:"true"`
	GoogleRedirectURL  string `split_words:"true"`
	GoogleAuthURL      string `split_words:"true"`
	GoogleTokenURL     string `split_words:"true"`
	GoogleUserURL      string `split_words:"true"`
	Port               int
	WebDirectory       string `split_words:"true"`
}

func DefaultConfig() *Config {
	return &Config{
		DatabaseMigrations: "file://services/app/migrations",
		DatabaseType:       "postgres",
		DatabaseURL:        "postgres://user:password@localhost:5432/db?sslmode=disable",
		GitHubRedirectURL:  "https://app-production-321806-o77vfhyfuq-ts.a.run.app/api/v1/auth/github/callback",
		GitHubAuthURL:      github.Endpoint.AuthURL,
		GitHubTokenURL:     github.Endpoint.TokenURL,
		GitHubUserURL:      "https://api.github.com/user",
		GitHubEmailsURL:    "https://api.github.com/user/emails",
		GoogleRedirectURL:  "https://app-production-321806-o77vfhyfuq-ts.a.run.app/api/v1/auth/google/callback",
		GoogleAuthURL:      google.Endpoint.AuthURL,
		GoogleTokenURL:     google.Endpoint.TokenURL,
		GoogleUserURL:      "https://www.googleapis.com/oauth2/v3/userinfo",
		Port:               4000,
		WebDirectory:       "./public",
	}
}

func LoadConfig() (*Config, error) {
	config := DefaultConfig()
	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}
	return config, nil
}
