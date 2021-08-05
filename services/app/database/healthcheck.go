package database

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

const maxRetryAttempts = 10

var ErrMaxRetryAttemptsExceeded = errors.New("max retry attempts exceeded")

func Healthcheck(db *sql.DB) error {
	return db.Ping()
}

func WaitUntilHealthy(db *sql.DB) error {
	log.Println("Waiting for healthy database...")
	for attempt := 1; attempt <= maxRetryAttempts; attempt++ {
		log.Printf("Attempt [%d/%d]\n", attempt, maxRetryAttempts)
		err := Healthcheck(db)
		if err == nil {
			log.Println("Database ready")
			return nil
		}
		log.Println("Error:", err)
		log.Println("Retrying in 5 seconds")
		time.Sleep(5 * time.Second)
	}
	return ErrMaxRetryAttemptsExceeded
}
