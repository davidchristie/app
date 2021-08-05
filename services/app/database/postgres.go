package database

import (
	"database/sql"
	"log"

	"github.com/davidchristie/app/services/app/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // file driver
	_ "github.com/lib/pq"                                // postgres driver
)

func NewPostgresConnection(config *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", config.DatabaseURL)
	if err != nil {
		return nil, err
	}
	if err = WaitUntilHealthy(db); err != nil {
		return nil, err
	}
	if err = MigratePostgres(config, db); err != nil {
		return nil, err
	}
	return db, nil
}

func MigratePostgres(config *config.Config, db *sql.DB) error {
	log.Println("Running database migrations...")
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(config.DatabaseMigrations, "postgres", driver)
	if err != nil {
		return err
	}
	m.Up()
	return nil
}
