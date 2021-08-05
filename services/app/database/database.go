package database

import (
	"database/sql"
	"errors"

	"github.com/davidchristie/app/services/app/config"
)

func NewConnection(config *config.Config) (*sql.DB, error) {
	switch config.DatabaseType {
	case "postgres":
		return NewPostgresConnection(config)
	default:
		return nil, errors.New("unsupported database type: " + config.DatabaseType)
	}
}
