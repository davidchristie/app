package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	DatabaseMigrations string `split_words:"true"`
	DatabaseType       string `split_words:"true"`
	DatabaseURL        string `split_words:"true"`
	Port               int
	WebDirectory       string `split_words:"true"`
}

func DefaultConfig() *Config {
	return &Config{
		DatabaseMigrations: "file://services/app/migrations",
		DatabaseType:       "postgres",
		DatabaseURL:        "postgres://user:password@localhost:5432/db?sslmode=disable",
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
