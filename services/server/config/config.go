package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Port         int
	WebDirectory string `split_words:"true"`
}

func DefaultConfig() *Config {
	return &Config{
		Port:         4000,
		WebDirectory: "./public",
	}
}

func LoadConfig() (*Config, error) {
	config := DefaultConfig()
	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}
	return config, nil
}
