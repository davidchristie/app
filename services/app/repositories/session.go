package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type SessionRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*Session, error)
	FindBySessionToken(ctx context.Context, sessionToken string) (*Session, error)
	Insert(ctx context.Context, session *Session) error
}

type Session struct {
	ID           uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ExpiresAt    time.Time
	SessionToken string
	UserID       uuid.UUID
}

type sessionRepository struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) SessionRepository {
	return &sessionRepository{
		db: db,
	}
}

func (s *sessionRepository) FindByID(ctx context.Context, id uuid.UUID) (*Session, error) {
	const query = `
		SELECT id, created_at, updated_at, expires_at, session_token, user_id FROM sessions
		WHERE id = $1
	`
	row := s.db.QueryRowContext(ctx, query, id)
	session := Session{}
	if err := row.Scan(
		&session.ID,
		&session.CreatedAt,
		&session.UpdatedAt,
		&session.ExpiresAt,
		&session.SessionToken,
		&session.UserID,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &session, nil
}

func (s *sessionRepository) FindBySessionToken(ctx context.Context, sessionToken string) (*Session, error) {
	const query = `
		SELECT id, created_at, updated_at, expires_at, session_token, user_id FROM sessions
		WHERE session_token = $1
	`
	row := s.db.QueryRowContext(ctx, query, sessionToken)
	session := Session{}
	if err := row.Scan(
		&session.ID,
		&session.CreatedAt,
		&session.UpdatedAt,
		&session.ExpiresAt,
		&session.SessionToken,
		&session.UserID,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &session, nil
}

func (s *sessionRepository) Insert(ctx context.Context, session *Session) error {
	const query = `
		INSERT INTO sessions (id, created_at, updated_at, expires_at, session_token, user_id)
		VALUES ($1, $2, $3, $4, $5, $6);
	`
	_, err := s.db.ExecContext(
		ctx,
		query,
		session.ID,
		session.CreatedAt,
		session.UpdatedAt,
		session.ExpiresAt,
		session.SessionToken,
		session.UserID,
	)
	if err != nil {
		return err
	}
	return nil
}
