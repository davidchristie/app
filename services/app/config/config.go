package config

import "github.com/kelseyhightower/envconfig"

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
	Port               int
	WebDirectory       string `split_words:"true"`
}

func DefaultConfig() *Config {
	return &Config{
		DatabaseMigrations: "file://services/app/migrations",
		DatabaseType:       "postgres",
		DatabaseURL:        "postgres://user:password@localhost:5432/db?sslmode=disable",
		GitHubRedirectURL:  "https://app-production-321806-o77vfhyfuq-ts.a.run.app/api/v1/auth/github/callback",
		GitHubAuthURL:      "https://github.com/login/oauth/authorize",
		GitHubTokenURL:     "https://github.com/login/oauth/access_token",
		GitHubUserURL:      "https://api.github.com/user",
		GitHubEmailsURL:    "https://api.github.com/user/emails",
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
